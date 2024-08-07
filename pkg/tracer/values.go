package tracer

var (
	// Loaded JS tracers for simulating various AiMiddleware methods using debug_traceCall.
	Loaded, _ = NewTracers()
)
