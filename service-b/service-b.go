package main

import (
	"fmt"
	"log"
	"net/http"

	"ping/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

const thisServiceName = "service-b"

func main() {
	tracer, closer := tracing.Init(thisServiceName)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		span := tracing.StartSpanFromRequest(tracer, r)
		defer span.Finish()

		w.Write([]byte(fmt.Sprintf("%s", thisServiceName)))
	})
	log.Printf("Listening on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
