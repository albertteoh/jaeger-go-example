all: start

.PHONY: start
start: service-a service-b
	@docker compose up --build --remove-orphans

.PHONY: stop
stop:
	@docker compose down
	docker rmi service-a service-b
