gen-proto:
	protoc -I=./proto --go_out=. \
	--go_opt=Mproto/product.proto=github.com/meokg456/ecommerce/proto/product \
 	--go_opt=Mproto/common.proto=github.com/meokg456/ecommerce/proto/common \
	--go-grpc_out=. \
	$(shell find ./proto -name '*.proto')