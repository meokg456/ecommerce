.PHONY: run test local-db lint db/migrate

run:
	air -c .air.toml

test:
	go clean -testcache
	go test -cover ./...

local-db:
	docker-compose --env-file ./.env -f ./tools/compose/docker-compose.yml down
	docker-compose --env-file ./.env -f ./tools/compose/docker-compose.yml up -d

lint:
	@(hash golangci-lint 2>/dev/null || \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(go env GOPATH)/bin v1.54.2)
	@golangci-lint run

db/migrate:
	go run ./cmd/migrate

seed-data:
	go run ./cmd/seeddata

kafka:
	go run ./cmd/kafka

unit-test:
	@mkdir coverage || true
	-go test -p 20 -race -v -coverprofile=coverage/coverage.txt.tmp -count=1 ./adapter/...
	-go test -p 20 -race -v -coverprofile=coverage/coverage.txt.tmp -count=1 ./domain/...
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage.txt
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage-2.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -func=coverage/coverage-2.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index-adapter.html
	@go tool cover -html=coverage/coverage-2.txt -o coverage/index-domain.html

gen-proto:
	protoc -I=./adapter/grpcserver/proto --go_out=adapter \
	--go-grpc_out=adapter \
	$(shell find ./adapter/grpcserver/proto -name '*.proto')

docker-build:
	docker build -t user-service:latest .

docker-db:
	docker run -d --name order-db  -p 5082:8000 amazon/dynamodb-local:1.19.0

docker-kafka:
	docker run -d -p 9085:9092 --name order-broker \
  	-e KAFKA_NODE_ID=1 \
  	-e KAFKA_PROCESS_ROLES=broker,controller \
  	-e KAFKA_LISTENERS=INTERNAL://0.0.0.0:9092,EXTERNAL://0.0.0.0:9085,CONTROLLER://localhost:9093 \
  	-e KAFKA_ADVERTISED_LISTENERS=EXTERNAL://localhost:9092,INTERNAL://localhost:9085 \
  	-e KAFKA_CONTROLLER_LISTENER_NAMES=CONTROLLER \
  	-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT \
  	-e KAFKA_CONTROLLER_QUORUM_VOTERS=1@localhost:9093 \
	-e KAFKA_INTER_BROKER_LISTENER_NAME=INTERNAL \
	-e KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS=0 \
	-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
	apache/kafka:4.0.0

docker-rm-kafka:
	docker container rm order-broker -f