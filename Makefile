GO = go

all: start

.PHONY: start
start: service-a service-b
	@docker compose up --build --remove-orphans

.PHONY: stop
stop:
	@docker compose down
	docker rmi service-a service-b

.PHONY: lint
lint:
	golangci-lint run

.PHONY: install-tools
install-tools:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
