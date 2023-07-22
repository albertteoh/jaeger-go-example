package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/albertteoh/jaeger-go-example/lib/ping"
	"github.com/albertteoh/jaeger-go-example/lib/tracing"
)

const thisServiceName = "service-a"

func main() {
	ctx := context.Background()
	tracer := tracing.Init(ctx, thisServiceName)

	outboundHostPort, ok := os.LookupEnv("OUTBOUND_HOST_PORT")
	if !ok {
		outboundHostPort = "localhost:8082"
	}

	http.HandleFunc("/ping", func(writer http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "/ping")
		defer span.End()

		response, err := ping.Ping(ctx, outboundHostPort, tracer)
		if err != nil {
			log.Fatalf("Error occurred on ping: %s", err)
		}
		if _, err = writer.Write([]byte(fmt.Sprintf("%s -> %s", thisServiceName, response))); err != nil {
			log.Fatalf("Error occurred on write: %s", err)
		}
	})
	log.Printf("Listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
