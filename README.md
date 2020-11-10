# Jaeger Go Instrumentation Example
Two simple Go microservices, exposing `/ping` endpoints, instrumented with Jaeger+OpenTracing.

# Getting Started
Firstly, startup jaeger all-in-one; which brings up the entire Jaeger frontend and backend components.

## Jaeger All-in-one
```
$ docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.20
```

## Start the services

In one terminal:

```
$ go run ./service-a/service-a.go
```

In a second terminal:

```
$ go run ./service-b/service-b.go
```

## Run the example

In a third terminal:

```
$ curl http://localhost:8081/ping
```

## Validate

Should see `service-a -> service-b` on STDOUT.

Go to http://localhost:16686/ and select `service-a` from the "Service" dropdown and click the "Find Traces" button.
