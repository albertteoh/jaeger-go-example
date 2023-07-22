package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// StartSpanFromRequest extracts the parent span context from the inbound HTTP request
// and starts a new child span if there is a parent span.
func StartSpanFromRequest(r *http.Request, tracer trace.Tracer) (context.Context, trace.Span) {
	p := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	ctx := p.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	return tracer.Start(ctx, "ping-receive")
}
