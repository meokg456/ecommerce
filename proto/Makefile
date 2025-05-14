gen-proto:
	protoc -I=./proto --go_out=. \
	--go-grpc_out=. \
	$(shell find ./proto -name '*.proto')