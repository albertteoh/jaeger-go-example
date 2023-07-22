package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/otel/propagation"

	"ping/lib/tracing"
)

const thisServiceName = "service-b"

func main() {
	ctx := context.Background()
	tracer := tracing.Init(ctx, thisServiceName)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("DEBUG: http req=%+v\n", r)

		p := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
		ctx = p.Extract(r.Context(), propagation.HeaderCarrier(r.Header))

		_, span := tracer.Start(ctx, "ping-receive")
		defer span.End()

		if _, err := w.Write([]byte(fmt.Sprintf("%s", thisServiceName))); err != nil {
			log.Fatalf("Error occurred on write: %s", err)
		}
	})
	log.Printf("Listening on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
