package main

import (
	"context"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/propagation"

	"github.com/albertteoh/jaeger-go-example/lib/tracing"
)

const thisServiceName = "service-b"

func main() {
	ctx := context.Background()
	tracer := tracing.Init(ctx, thisServiceName)

	http.HandleFunc("/ping", func(writer http.ResponseWriter, r *http.Request) {
		p := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
		ctx = p.Extract(r.Context(), propagation.HeaderCarrier(r.Header))

		_, span := tracer.Start(ctx, "ping-receive")
		defer span.End()

		if _, err := writer.Write([]byte(thisServiceName)); err != nil {
			log.Fatalf("Error occurred on write: %s", err)
		}
	})
	log.Printf("Listening on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
