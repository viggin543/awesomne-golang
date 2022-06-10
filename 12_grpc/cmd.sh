go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc cart.proto "-I" . \
      "--go_opt=paths=source_relative"\
      "--go_out=gen" "--go-grpc_out=gen"\
      "--go-grpc_opt=require_unimplemented_servers=false"\
      "--go-grpc_opt=paths=source_relative"


protoc -I  . "--grpc-gateway_out=logtostderr=true,paths=source_relative:gen"  cart.proto

protoc -I  . "--swagger_out=logtostderr=true:gen" cart.proto