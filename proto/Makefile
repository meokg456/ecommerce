gen-proto:
	protoc --go_out=. \
	--go-grpc_out=. \
	--proto_path=. \
	--go_opt=module=github.com/meokg456/ecommerce/proto \
	--go-grpc_opt=module=github.com/meokg456/ecommerce/proto \
	$(shell find ./proto -name '*.proto')