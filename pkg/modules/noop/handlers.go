// Package noop implements basic no-operation modules which are used by default for both Client and Bundler.
package noop

import "github.com/DAO-Metaplayer/aiops-bundler/pkg/modules"

// BatchHandler takes a BatchHandlerCtx and returns nil error.
func BatchHandler(ctx *modules.BatchHandlerCtx) error {
	return nil
}

// AiOpHandler takes a AiOpHandlerCtx and returns nil error.
func AiOpHandler(ctx *modules.AiOpHandlerCtx) error {
	return nil
}
