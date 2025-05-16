module = github.com/meokg456/ecommerce/proto

gen-proto:
	protoc --go_out=. \
	--go-grpc_out=. \
	--proto_path=. \
	--go_opt=module=$(module) \
	--go-grpc_opt=module=$(module) \
	$(shell find ./proto -name '*.proto')