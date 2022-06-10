package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	_ "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	cart "github.com/viggin543/awesomne-golang/12_grpc/gen"
	grpc "google.golang.org/grpc"
	"net/http"
)

type GrpcServer struct {
	grpcPort   int
	restPort   int
	tracer     opentracing.Tracer
	grpcServer *grpc.Server
	restServer *http.ServeMux
}

func (this *GrpcServer) UpsertCart(ctx context.Context, c *cart.Cart) (*cart.Cart, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	impl := GrpcServer{}
	NewGrpcServer(&impl)
	NewRestGateway(10000, grpc.WithInsecure())
}

func NewGrpcServer(impl cart.CartSvcServer) {
	unaryInterceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_recovery.UnaryServerInterceptor()))
	streamInterceptors := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_recovery.StreamServerInterceptor()))
	server := grpc.NewServer(unaryInterceptors, streamInterceptors)
	server.RegisterService(&cart.CartSvc_ServiceDesc, impl)
}

func NewRestGateway(grpcDestPort int, opts ...grpc.DialOption) {
	ctx := context.Background()

	grpcMux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(anyHeaderMatcher),
		runtime.WithIncomingHeaderMatcher(anyHeaderMatcher),
	)

	grpcAddress := fmt.Sprintf("localhost:%d", grpcDestPort)
	err := cart.RegisterCartSvcHandlerFromEndpoint(ctx, grpcMux, grpcAddress, opts)
	if err != nil {
		panic(err)
	}

}

func anyHeaderMatcher(header string) (string, bool) {
	if header == "Connection" {
		return header, false
	}
	return header, true
}
