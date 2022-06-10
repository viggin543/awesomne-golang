package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	_ "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	_ "github.com/viggin543/awesomne-golang/12_grpc"
	grpc "google.golang.org/grpc"
)

func main() {
	var err error
	var grpcOptions []grpc.ServerOption
	unaryOpts := []grpc.UnaryServerInterceptor{
		grpc_recovery.UnaryServerInterceptor(),
	}

	unaryInterceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryOpts...))

	streamOpts := []grpc.StreamServerInterceptor{
		grpc_recovery.StreamServerInterceptor(),
	}
	streamInterceptors := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamOpts...))

	server := grpc.NewServer(unaryInterceptors, streamInterceptors)
	server.RegisterService(&CatalogService_ServiceDesc, impl)
	grpc_prometheus.Register(server)
	grpc_prometheus.EnableHandlingTimeHistogram()

	var restOptions = []grpc.DialOption{grpc.WithInsecure()}
	this.restServer, _, err = catalog.NewRestGateway(restPort, grpcPort, restOptions...)
	if err != nil {
		return nil, err
	}
}
