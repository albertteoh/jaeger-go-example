package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

// StartSpanFromRequest extracts the parent span context from the inbound HTTP request
// and starts a new child span if there is a parent span.
func StartSpanFromRequest(tracer opentracing.Tracer, r *http.Request) opentracing.Span {
	spanCtx, _ := Extract(tracer, r)
	return tracer.StartSpan("ping-receive", ext.RPCServerOption(spanCtx))
}
