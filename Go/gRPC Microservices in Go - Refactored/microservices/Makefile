# gRPC Microservices in Go: Order + Payment + Shipping
# Sync: Order → Payment (gRPC). Async: Order ⇄ Shipping (Kafka).
# Run `make` or `make help` to list the targets.

SHELL := /usr/bin/env bash
.DEFAULT_GOAL := help

SERVICES := order payment shipping
MODULES  := proto $(SERVICES) e2e tools

# Pinned tools, run through `go run`, so nothing has to be installed globally.
BUF     := go run github.com/bufbuild/buf/cmd/buf@v1.75.0
MOCKERY := go run github.com/vektra/mockery/v3@v3.9.0
GRPCURL := go run github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.9.4

# `buf breaking` compares against this Git ref (subdir = where buf.yaml lives).
BREAKING_AGAINST ?= .git\#branch=main,subdir=proto

COMPOSE     := docker compose
COMPOSE_TLS := $(COMPOSE) -f docker-compose.yml -f docker-compose.tls.yml

ORDER_ADDR    ?= localhost:8080
SHIPPING_ADDR ?= localhost:8082
KAFKA_BIN     := $(COMPOSE) exec kafka /opt/kafka/bin
KAFKA         := --bootstrap-server localhost:9092
ORDER_ID   ?= 1
MYSQL_PORT ?= 3306

##@ Help
.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: make \033[36m<target>\033[0m\n"} \
	/^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } \
	/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

##@ Protobuf (Buf)
.PHONY: proto proto-lint proto-format proto-breaking proto-generate
proto: proto-format proto-lint proto-generate ## Format, lint, and generate the Go stubs

proto-lint: ## Check the .proto files against buf.yaml lint rules
	cd proto && $(BUF) lint

proto-format: ## Rewrite the .proto files in the standard format
	cd proto && $(BUF) format -w

proto-breaking: ## Fail on breaking changes vs BREAKING_AGAINST (default: main branch)
	cd proto && $(BUF) breaking --against '$(BREAKING_AGAINST)'

proto-generate: ## Generate gen/go/** with the pinned remote plugins
	cd proto && $(BUF) generate && GOWORK=off go mod tidy

##@ Go
.PHONY: tidy mocks build vet test test-integration cover
tidy: ## go mod tidy in every module
	@for m in $(MODULES); do (cd $$m && GOWORK=off go mod tidy) || exit 1; done

mocks: ## Regenerate testify mocks for the ports (mockery v3)
	@for s in $(SERVICES); do (cd $$s && $(MOCKERY)) || exit 1; done

build: ## Compile both services into ./bin
	@mkdir -p bin
	@for s in $(SERVICES); do (cd $$s && go build -o ../bin/$$s ./cmd) || exit 1; done

vet: ## go vet every module
	@for m in $(MODULES); do (cd $$m && go vet ./...) || exit 1; done

test: ## Unit tests only (fast, no Docker)
	@for s in $(SERVICES); do (cd $$s && go test -short -race ./...) || exit 1; done

test-integration: ## Unit + MySQL integration tests (Testcontainers, needs Docker)
	@for s in $(SERVICES); do (cd $$s && go test -race -count=1 ./...) || exit 1; done

cover: ## Unit-test coverage report per service (coverage-<service>.html)
	@for s in $(SERVICES); do \
		(cd $$s && go test -short -coverprofile=../coverage-$$s.out ./... && \
		 go tool cover -html=../coverage-$$s.out -o ../coverage-$$s.html) || exit 1; \
	done

##@ Run locally (services on the host, MySQL + Kafka in Docker)
.PHONY: db-up run-payment run-order run-shipping
db-up: ## Start MySQL (MYSQL_PORT) and Kafka (localhost:9092) with topics
	$(COMPOSE) up -d --wait mysql kafka
	$(COMPOSE) run --rm kafka-init

run-payment: ## Run Payment on :3001
	cd payment && ENV=development APPLICATION_PORT=3001 \
	DATA_SOURCE_URL="root:s3cr3t@tcp(127.0.0.1:$(MYSQL_PORT))/payments?charset=utf8mb4&parseTime=True&loc=Local" \
	go run ./cmd

run-order: ## Run Order on :3000, calling Payment on :3001
	cd order && ENV=development APPLICATION_PORT=3000 PAYMENT_SERVICE_URL=localhost:3001 \
	KAFKA_BROKERS=localhost:9092 \
	DATA_SOURCE_URL="root:s3cr3t@tcp(127.0.0.1:$(MYSQL_PORT))/orders?charset=utf8mb4&parseTime=True&loc=Local" \
	go run ./cmd

run-shipping: ## Run Shipping on :3002
	cd shipping && ENV=development APPLICATION_PORT=3002 KAFKA_BROKERS=localhost:9092 \
	DATA_SOURCE_URL="root:s3cr3t@tcp(127.0.0.1:$(MYSQL_PORT))/shipments?charset=utf8mb4&parseTime=True&loc=Local" \
	go run ./cmd

##@ Docker Compose (whole stack)
.PHONY: docker-build up up-tls down logs ps certs
docker-build: ## Build both service images
	$(COMPOSE) build

up: ## Build and start MySQL, Kafka, Payment, Order, Shipping, Jaeger
	$(COMPOSE) up -d --build --wait

up-tls: certs ## Same as `up`, with mTLS between Order and Payment
	$(COMPOSE_TLS) up -d --build --wait

down: ## Stop the stack and delete its data
	$(COMPOSE) down -v --remove-orphans

logs: ## Follow the service logs
	$(COMPOSE) logs -f order payment shipping

ps: ## Show container status
	$(COMPOSE) ps

certs: ## Create a dev CA and mTLS certificates in ./certs
	@[ -f certs/ca-cert.pem ] || ./scripts/gen-certs.sh certs

##@ Call the API (stack must be running)
.PHONY: list create-order get-order get-shipment bad-order buf-curl-create
list: ## List services via reflection
	$(GRPCURL) -plaintext $(ORDER_ADDR) list

create-order: ## Place an order with grpcurl
	$(GRPCURL) -plaintext -d '{"user_id": 23, "order_items": [{"product_code": "CAM123", "unit_price": 12.5, "quantity": 2}]}' \
		$(ORDER_ADDR) order.v1.OrderService/Create

get-order: ## Read an order (ORDER_ID=1)
	$(GRPCURL) -plaintext -d '{"order_id": $(ORDER_ID)}' $(ORDER_ADDR) order.v1.OrderService/Get

get-shipment: ## Read the shipment Shipping created for ORDER_ID (async result)
	$(GRPCURL) -plaintext -d '{"order_id": $(ORDER_ID)}' $(SHIPPING_ADDR) shipping.v1.ShippingService/Get

bad-order: ## Send an invalid order to see InvalidArgument + field violations
	-$(GRPCURL) -plaintext -d '{"user_id": 0, "order_items": [{"product_code": "", "unit_price": 0, "quantity": 0}]}' \
		$(ORDER_ADDR) order.v1.OrderService/Create

buf-curl-create: ## Place an order with buf curl (uses reflection)
	$(BUF) curl --protocol grpc --http2-prior-knowledge \
		--data '{"user_id": 23, "order_items": [{"product_code": "CAM123", "unit_price": 12.5, "quantity": 2}]}' \
		http://$(ORDER_ADDR)/order.v1.OrderService/Create

##@ Kafka (stack must be running)
.PHONY: topics events dlq lag poison
topics: ## Describe the topics: partitions, leaders, replicas
	$(KAFKA_BIN)/kafka-topics.sh $(KAFKA) --describe --exclude-internal

events: ## Print the events in "orders" and "shipments" as JSON
	cd tools && go run ./eventtail -topic orders && go run ./eventtail -topic shipments

dlq: ## Print the dead-letter topics
	cd tools && go run ./eventtail -topic shipping.dlq && go run ./eventtail -topic order.dlq

lag: ## Show every consumer group's committed offsets and lag
	$(KAFKA_BIN)/kafka-consumer-groups.sh $(KAFKA) --describe --all-groups

poison: ## Send an undecodable record to "orders" (it should land in shipping.dlq)
	printf '999:not-protobuf\n' | $(COMPOSE) exec -T kafka /opt/kafka/bin/kafka-console-producer.sh $(KAFKA) \
		--topic orders --reader-property parse.key=true --reader-property key.separator=:

##@ Tests against the whole stack
.PHONY: e2e
e2e: docker-build ## End-to-end test: Testcontainers starts the Compose stack
	cd e2e && go test -count=1 -v ./...

##@ Kubernetes (minikube)
.PHONY: k8s-load k8s-deploy k8s-delete
k8s-load: docker-build ## Copy the local images into minikube
	minikube image load microservices/order:dev
	minikube image load microservices/payment:dev
	minikube image load microservices/shipping:dev

k8s-deploy: k8s-load ## Apply deploy/k8s with kustomize
	kubectl apply -k deploy/k8s

k8s-delete: ## Remove everything deploy/k8s created
	kubectl delete -k deploy/k8s --ignore-not-found

##@ Clean
.PHONY: clean
clean: ## Remove build output, coverage files, and certificates
	rm -rf bin certs coverage-*.out coverage-*.html
