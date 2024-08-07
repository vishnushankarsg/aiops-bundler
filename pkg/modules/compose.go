package modules

// ComposeBatchHandlerFunc combines many BatchHandlers into one.
func ComposeBatchHandlerFunc(fns ...BatchHandlerFunc) BatchHandlerFunc {
	return func(ctx *BatchHandlerCtx) error {
		for _, fn := range fns {
			err := fn(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// ComposeAiOpHandlerFunc combines many AiOpHandlers into one.
func ComposeAiOpHandlerFunc(fns ...AiOpHandlerFunc) AiOpHandlerFunc {
	return func(ctx *AiOpHandlerCtx) error {
		for _, fn := range fns {
			err := fn(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
