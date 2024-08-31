package start

import (
	"context"
	"fmt"
	"log"
	"net/http"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/metachris/flashbotsrpc"
	"gitlab.com/quantum-warriors/aiops-bundler/internal/config"
	"gitlab.com/quantum-warriors/aiops-bundler/internal/logger"
	"gitlab.com/quantum-warriors/aiops-bundler/internal/o11y"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/aimiddleware/stake"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/altmempools"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/bundler"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/client"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/gas"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/jsonrpc"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/mempool"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/modules/batch"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/modules/builder"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/modules/checks"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/modules/entities"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/modules/expire"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/modules/gasprice"
	"gitlab.com/quantum-warriors/aiops-bundler/pkg/signer"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
)

func SearcherMode() {
	conf := config.GetValues()

	logr := logger.NewZeroLogr().
		WithName("aiops_bundler").
		WithValues("bundler_mode", "searcher")

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

	fb := flashbotsrpc.NewBuilderBroadcastRPC(conf.EthBuilderUrls)

	chain, err := eth.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if !builder.CompatibleChainIDs.Contains(chain.Uint64()) {
		log.Fatalf(
			"error: network with chainID %d is not compatible with the Block Builder API.",
			chain.Uint64(),
		)
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

	// TODO: Create separate go-routine for tracking transactions sent to the block builder.
	builder := builder.New(eoa, eth, fb, beneficiary, conf.BlocksInTheFuture)

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
		// TODO: add p2p propagation module
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
		builder.SendAiOperation(),
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
