generate:
	echo "generating grpc messages and service"
	protoc cart.proto "-I" . \
          "--go_opt=paths=source_relative"\
          "--go_out=gen" "--go-grpc_out=gen"\
          "--go-grpc_opt=require_unimplemented_servers=false"\
          "--go-grpc_opt=paths=source_relative"
	echo "generating rest proxy"
	protoc -I  . "--grpc-gateway_out=logtostderr=true,paths=source_relative:gen"  cart.proto
	echo "generating rest swagger definitions"
	protoc -I  . "--swagger_out=logtostderr=true:gen" cart.proto

build:
	docker build -t viggin543/protoc .
	docker run --rm -v $(shell pwd):/work viggin543/protoc
