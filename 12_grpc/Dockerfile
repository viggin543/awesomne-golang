ARG GO_VERSION=1.16
FROM golang:${GO_VERSION}-alpine

ENV PROTOBUF_TAG="v3.12.2"

RUN apk update
RUN apk add --update libstdc++
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN apk add --no-cache bash git vim make protobuf protobuf-dev
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/favadi/protoc-go-inject-tag
RUN go install github.com/vektra/mockery/v2@latest

WORKDIR /work
RUN protoc --version
ENTRYPOINT ["make","generate"]
