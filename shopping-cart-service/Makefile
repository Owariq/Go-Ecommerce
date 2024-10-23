
install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-auth-api

generate-product-api:
	protoc --proto_path api\ \
	--go_out=pkg\cart_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg\cart_v1 --go-grpc_opt=paths=source_relative \
	api\cart.proto
