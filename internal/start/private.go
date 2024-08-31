package start

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AO-Metaplayer/aiops-bundler/internal/config"
	"github.com/AO-Metaplayer/aiops-bundler/internal/logger"
	"github.com/AO-Metaplayer/aiops-bundler/internal/o11y"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aimiddleware/stake"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/altmempools"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/bundler"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/client"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/gas"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/jsonrpc"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/mempool"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/batch"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/checks"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/entities"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/expire"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/gasprice"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/relay"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/signer"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

func PrivateMode() {
	conf := config.GetValues()

	logr := logger.NewZeroLogr().
		WithName("aiops_bundler").
		WithValues("bundler_mode", "private")

	eoa, err := signer.New(conf.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	beneficiary := common.HexToAddress(conf.Beneficiary)

	db, err := badger.Open(badger.DefaultOptions(conf.DataDirectory))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	runDBGarbageCollection(db)

	rpc, err := rpc.Dial(conf.EthClientUrl)
	if err != nil {
		log.Fatal(err)
	}

	eth := ethclient.NewClient(rpc)

	chain, err := eth.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if o11y.IsEnabled(conf.OTELServiceName) {
		o11yOpts := &o11y.Opts{
			ServiceName:     conf.OTELServiceName,
			CollectorHeader: conf.OTELCollectorHeaders,
			CollectorUrl:    conf.OTELCollectorUrl,
			InsecureMode:    conf.OTELInsecureMode,

			ChainID: chain,
			Address: eoa.Address,
		}

		tracerCleanup := o11y.InitTracer(o11yOpts)
		defer tracerCleanup()

		metricsCleanup := o11y.InitMetrics(o11yOpts)
		defer metricsCleanup()
	}

	ov := gas.NewDefaultOverhead()
	if conf.IsArbStackNetwork || config.ArbStackChains.Contains(chain.Uint64()) {
		ov.SetCalcPreVerificationGasFunc(gas.CalcArbitrumPVGWithEthClient(rpc, conf.SupportedAiMiddlewares[0]))
		ov.SetPreVerificationGasBufferFactor(16)
	}

	if conf.IsOpStackNetwork || config.OpStackChains.Contains(chain.Uint64()) {
		ov.SetCalcPreVerificationGasFunc(
			gas.CalcOptimismPVGWithEthClient(rpc, chain, conf.SupportedAiMiddlewares[0]),
		)
		ov.SetPreVerificationGasBufferFactor(1)
	}

	mem, err := mempool.New(db)
	if err != nil {
		log.Fatal(err)
	}

	alt, err := altmempools.NewFromIPFS(chain, conf.AltMempoolIPFSGateway, conf.AltMempoolIds)
	if err != nil {
		log.Fatal(err)
	}

	check := checks.New(
		db,
		rpc,
		ov,
		alt,
		conf.MaxVerificationGas,
		conf.MaxBatchGasLimit,
		conf.IsRIP7212Supported,
		conf.NativeBundlerCollectorTracer,
		conf.ReputationConstants,
	)

	exp := expire.New(conf.MaxOpTTL)

	relayer := relay.New(eoa, eth, chain, beneficiary, logr)

	rep := entities.New(db, eth, conf.ReputationConstants)

	// Init Client
	c := client.New(mem, ov, chain, conf.SupportedAiMiddlewares, conf.OpLookupLimit)
	c.SetGetAiOpReceiptFunc(client.GetAiOpReceiptWithEthClient(eth))
	c.SetGetGasPricesFunc(client.GetGasPricesWithEthClient(eth))
	c.SetGetGasEstimateFunc(
		client.GetGasEstimateWithEthClient(
			rpc,
			ov,
			chain,
			conf.MaxBatchGasLimit,
			conf.NativeBundlerExecutorTracer,
		),
	)
	c.SetGetAiOpByHashFunc(client.GetAiOpByHashWithEthClient(eth))
	c.SetGetStakeFunc(stake.GetStakeWithEthClient(eth))
	c.UseLogger(logr)
	c.UseModules(
		rep.CheckStatus(),
		rep.ValidateOpLimit(),
		check.ValidateOpValues(),
		check.SimulateOp(),
		rep.IncOpsSeen(),
	)

	// Init Bundler
	b := bundler.New(mem, chain, conf.SupportedAiMiddlewares)
	b.SetGetBaseFeeFunc(gasprice.GetBaseFeeWithEthClient(eth))
	b.SetGetGasTipFunc(gasprice.GetGasTipWithEthClient(eth))
	b.SetGetLegacyGasPriceFunc(gasprice.GetLegacyGasPriceWithEthClient(eth))
	b.UseLogger(logr)
	if err := b.AiMeter(otel.GetMeterProvider().Meter("bundler")); err != nil {
		log.Fatal(err)
	}
	b.UseModules(
		exp.DropExpired(),
		gasprice.SortByGasPrice(),
		gasprice.FilterUnderpriced(),
		batch.SortByNonce(),
		batch.MaintainGasLimit(conf.MaxBatchGasLimit),
		check.CodeHashes(),
		check.PaymasterDeposit(),
		check.SimulateBatch(),
		relayer.SendAiOperation(),
		rep.IncOpsIncluded(),
		check.Clean(),
	)
	if err := b.Run(); err != nil {
		log.Fatal(err)
	}

	// init Debug
	var d *client.Debug
	if conf.DebugMode {
		d = client.NewDebug(eoa, eth, mem, rep, b, chain, conf.SupportedAiMiddlewares[0], beneficiary)
		b.SetMaxBatch(1)
		relayer.SetWaitTimeout(0)
	}

	// Init HTTP server
	gin.SetMode(conf.GinMode)
	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}
	if o11y.IsEnabled(conf.OTELServiceName) {
		r.Use(otelgin.Middleware(conf.OTELServiceName))
	}
	r.Use(
		cors.Default(),
		logger.WithLogr(logr),
		gin.Recovery(),
	)
	r.GET("/ping", func(g *gin.Context) {
		g.Status(http.StatusOK)
	})
	handlers := []gin.HandlerFunc{
		jsonrpc.Controller(client.NewRpcAdapter(c, d)),
		jsonrpc.WithOTELTracerAttributes(),
	}
	r.POST("/", handlers...)
	r.POST("/rpc", handlers...)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		log.Fatal(err)
	}
}
