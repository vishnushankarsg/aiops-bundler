// Package modules provides standard interfaces for extending the Client and Bundler packages with
// middleware.
package modules

// BatchHandlerFunc is an interface to support modular processing of AiOperation batches by the Bundler.
type BatchHandlerFunc func(ctx *BatchHandlerCtx) error

// OpHandlerFunc is an interface to support modular processing of single AiOperations by the Client.
type AiOpHandlerFunc func(ctx *AiOpHandlerCtx) error
