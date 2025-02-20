// Package bundler provides the mediator for processing outgoing AiOperation batches to the AiMiddleware.
package bundler

import (
	"context"
	"math/big"
	"time"

	"github.com/AO-Metaplayer/aiops-bundler/internal/logger"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/aiop"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/mempool"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/gasprice"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/noop"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// Bundler controls the end to end process of creating a batch of AiOperations from the mempool and sending
// it to the AiMiddleware.
type Bundler struct {
	mempool                *mempool.Mempool
	chainID                *big.Int
	supportedAiMiddlewares []common.Address
	batchHandler           modules.BatchHandlerFunc
	logger                 logr.Logger
	meter                  metric.Meter
	isRunning              bool
	done                   chan bool
	stop                   func()
	maxBatch               int
	gbf                    gasprice.GetBaseFeeFunc
	ggt                    gasprice.GetGasTipFunc
	ggp                    gasprice.GetLegacyGasPriceFunc
}

// New initializes a new EIP-4337 bundler which can be extended with modules for validating batches and
// excluding AiOperations that should not be sent to the AiMiddleware and/or dropped from the mempool.
func New(mempool *mempool.Mempool, chainID *big.Int, supportedAiMiddlewares []common.Address) *Bundler {
	return &Bundler{
		mempool:                mempool,
		chainID:                chainID,
		supportedAiMiddlewares: supportedAiMiddlewares,
		batchHandler:           noop.BatchHandler,
		logger:                 logger.NewZeroLogr().WithName("bundler"),
		meter:                  otel.GetMeterProvider().Meter("bundler"),
		isRunning:              false,
		done:                   make(chan bool),
		stop:                   func() {},
		maxBatch:               0,
		gbf:                    gasprice.NoopGetBaseFeeFunc(),
		ggt:                    gasprice.NoopGetGasTipFunc(),
		ggp:                    gasprice.NoopGetLegacyGasPriceFunc(),
	}
}

// SetMaxBatch defines the max number of AiOperations per bundle. The default value is 0 (i.e. unlimited).
func (i *Bundler) SetMaxBatch(max int) {
	i.maxBatch = max
}

// SetGetBaseFeeFunc defines the function used to retrieve an estimate for basefee during each bundler run.
func (i *Bundler) SetGetBaseFeeFunc(gbf gasprice.GetBaseFeeFunc) {
	i.gbf = gbf
}

// SetGetGasTipFunc defines the function used to retrieve an estimate for gas tip during each bundler run.
func (i *Bundler) SetGetGasTipFunc(ggt gasprice.GetGasTipFunc) {
	i.ggt = ggt
}

// SetGetLegacyGasPriceFunc defines the function used to retrieve an estimate for gas price during each
// bundler run.
func (i *Bundler) SetGetLegacyGasPriceFunc(ggp gasprice.GetLegacyGasPriceFunc) {
	i.ggp = ggp
}

// UseLogger defines the logger object used by the Bundler instance based on the go-logr/logr interface.
func (i *Bundler) UseLogger(logger logr.Logger) {
	i.logger = logger.WithName("bundler")
}

// AiMeter defines an opentelemetry meter object used by the Bundler instance to capture metrics during each
// run.
func (i *Bundler) AiMeter(meter metric.Meter) error {
	i.meter = meter
	_, err := i.meter.Int64ObservableGauge(
		"bundler_mempool_size",
		metric.WithInt64Callback(func(ctx context.Context, io metric.Int64Observer) error {
			size := 0
			for _, ep := range i.supportedAiMiddlewares {
				batch, err := i.mempool.Dump(ep)
				if err != nil {
					return err
				}
				size += len(batch)
			}
			io.Observe(int64(size))
			return nil
		}),
	)
	return err
}

// UseModules defines the BatchHandlers to process batches after it has gone through the standard checks.
func (i *Bundler) UseModules(handlers ...modules.BatchHandlerFunc) {
	i.batchHandler = modules.ComposeBatchHandlerFunc(handlers...)
}

// Process will create a batch from the mempool and send it through to the AiMiddleware.
func (i *Bundler) Process(ep common.Address) (*modules.BatchHandlerCtx, error) {
	// Init logger
	start := time.Now()
	l := i.logger.
		WithName("run").
		WithValues("aimiddleware", ep.String()).
		WithValues("chain_id", i.chainID.String())

	// Get all pending aiOps from the mempool. This will be in FIFO order. Downstream modules should sort it
	// based on more specific strategies.
	batch, err := i.mempool.Dump(ep)
	if err != nil {
		l.Error(err, "bundler run error")
		return nil, err
	}
	if len(batch) == 0 {
		return nil, nil
	}
	batch = adjustBatchSize(i.maxBatch, batch)

	// Get current block basefee
	bf, err := i.gbf()
	if err != nil {
		l.Error(err, "bundler run error")
		return nil, err
	}

	// Get suggested gas tip
	var gt *big.Int
	if bf != nil {
		gt, err = i.ggt()
		if err != nil {
			l.Error(err, "bundler run error")
			return nil, err
		}
	}

	// Get suggested gas price (for networks that don't support EIP-1559)
	gp, err := i.ggp()
	if err != nil {
		l.Error(err, "bundler run error")
		return nil, err
	}

	// Create context and execute modules.
	ctx := modules.NewBatchHandlerContext(batch, ep, i.chainID, bf, gt, gp)
	if err := i.batchHandler(ctx); err != nil {
		l.Error(err, "bundler run error")
		return nil, err
	}

	// Remove aiOps that remain in the context from mempool.
	rmOps := append([]*aiop.AiOperation{}, ctx.Batch...)
	dh := []string{}
	dr := []string{}
	for _, item := range ctx.PendingRemoval {
		rmOps = append(rmOps, item.Op)
		dh = append(dh, item.Op.GetAiOpHash(ep, i.chainID).String())
		dr = append(dr, item.Reason)
	}
	if err := i.mempool.RemoveOps(ep, rmOps...); err != nil {
		l.Error(err, "bundler run error")
		return nil, err
	}

	// Update logs for the current run.
	bat := []string{}
	for _, op := range ctx.Batch {
		bat = append(bat, op.GetAiOpHash(ep, i.chainID).String())
	}
	l = l.WithValues("batch_aiop_hashes", bat)
	l = l.WithValues("dropped_aiop_hashes", dh)
	l = l.WithValues("dropped_aiop_reasons", dr)

	for k, v := range ctx.Data {
		l = l.WithValues(k, v)
	}
	l = l.WithValues("duration", time.Since(start))
	l.Info("bundler run ok")
	return ctx, nil
}

// Run starts a goroutine that will continuously process batches from the mempool.
func (i *Bundler) Run() error {
	if i.isRunning {
		return nil
	}

	ticker := time.NewTicker(1 * time.Second)
	go func(i *Bundler) {
		for {
			select {
			case <-i.done:
				return
			case <-ticker.C:
				for _, ep := range i.supportedAiMiddlewares {
					_, err := i.Process(ep)
					if err != nil {
						// Already logged.
						continue
					}
				}
			}
		}
	}(i)

	i.isRunning = true
	i.stop = ticker.Stop
	return nil
}

// Stop signals the bundler to stop continuously processing batches from the mempool.
func (i *Bundler) Stop() {
	if !i.isRunning {
		return
	}

	i.isRunning = false
	i.stop()
	i.done <- true
}
