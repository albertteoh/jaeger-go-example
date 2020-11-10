package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"ping/lib/ping"
	"ping/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

const thisServiceName = "service-a"

func main() {
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		span := tracing.StartSpanFromRequest(tracer, r)
		defer span.Finish()

		ctx := opentracing.ContextWithSpan(context.Background(), span)
		response, err := ping.Ping(ctx, "localhost:8082")
		if err != nil {
			log.Fatalf("Error occurred: %s", err)
		}
		w.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, response)))
	})
	log.Printf("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
