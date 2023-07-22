package ping

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	libhttp "ping/lib/http"
)

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.
func Ping(ctx context.Context, hostPort string, tracer trace.Tracer) (string, error) {
	ctx, span := tracer.Start(ctx, "ping-send")
	defer span.End()

	url := fmt.Sprintf("http://%s/ping", hostPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	propagator := propagation.TraceContext{}
	fmt.Printf("DEBUG: propagator=%+v\n", propagator)

	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	fmt.Printf("DEBUG: http req=%+v\n", req)
	return libhttp.Do(req)
}
