# build stage
FROM golang as builder

ENV GO111MODULE=on
ENV PORT 8082
ENV APP service-b

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ${APP}/ ${APP}
COPY lib/ lib

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${APP} ./${APP}

# final stage
FROM scratch
COPY --from=builder /out/${APP} /app/

EXPOSE ${PORT}
ENTRYPOINT ["/app/service-b"]
