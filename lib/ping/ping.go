package ping

import (
	"context"
	"fmt"
	"net/http"

	libhttp "ping/lib/http"
	"ping/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.
func Ping(ctx context.Context, hostPort string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ping-send")
	defer span.Finish()

	url := fmt.Sprintf("http://%s/ping", hostPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if err := tracing.Inject(span, req); err != nil {
		return "", err
	}
	return libhttp.Do(req)
}
