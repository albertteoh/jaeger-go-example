package ping

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	libhttp "github.com/albertteoh/jaeger-go-example/lib/http"
)

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.
func Ping(ctx context.Context, hostPort string, tracer trace.Tracer) (string, error) {
	ctx, span := tracer.Start(ctx, "ping-send")
	defer span.End()

	url := fmt.Sprintf("http://%s/ping", hostPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed GET request to %s: %w", url, err)
	}

	propagator := propagation.TraceContext{}

	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	respBody, err := libhttp.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed http request: %w", err)
	}

	return respBody, nil
}
