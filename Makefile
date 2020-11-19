all: service-a service-b start

.PHONY: service-a
service-a:
	@docker build -t service-a -f service-a/Dockerfile .

.PHONY: service-b
service-b:
	@docker build -t service-b -f service-b/Dockerfile .

.PHONY: start
start:
	@docker-compose up -d

.PHONY: stop
stop:
	@docker-compose down
