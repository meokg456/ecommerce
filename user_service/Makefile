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

load-data:
	go run ./cmd/loaddata

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

docker:
	docker build -t user-service:latest .

docker-db:
	docker run -p 5431:5432 --name user-db -e POSTGRES_PASSWORD=root -d postgres:17.5-alpine3.21