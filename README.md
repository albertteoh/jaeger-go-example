# Jaeger Go Instrumentation Example
Two simple Go microservices where `service-a` calls `service-b`. Both services expose a `/ping` endpoint, instrumented with Jaeger+OpenTracing.

# Getting Started

## Start the example

Starts up the Jaeger all-in-one container, along with our example microservices.
```
$ make start
```

## Run the example

Hit `service-a`'s endpoint to trigger the trace.
```
$ curl -w '\n' http://localhost:8081/ping
```

## Validate

Should see `service-a -> service-b` on STDOUT.

Go to http://localhost:16686/ and select `service-a` from the "Service" dropdown and click the "Find Traces" button.

## Stop the example

Stop and remove containers.

```
$ make stop
```
