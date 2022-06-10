package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	_ "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	cart "github.com/viggin543/awesomne-golang/12_grpc/gen"
	grpc "google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

type GrpcServer struct {
	grpcPort   int
	restPort   int
	grpcServer *grpc.Server
	restServer *http.ServeMux
}

func main() {
	impl := GrpcServer{grpcPort: 10000, restPort: 8080}
	impl.grpcServer = NewGrpcServer(&impl)
	impl.restServer = NewRestGateway(impl.grpcPort, grpc.WithInsecure())
	err := impl.Start()
	if err != nil {
		panic(err)
	}
}

func (this *GrpcServer) UpsertCart(ctx context.Context, c *cart.Cart) (*cart.Cart, error) {
	log.Println("this is so cool...")
	c.TotalCents += 1
	return c, nil
}

func NewGrpcServer(impl cart.CartSvcServer) *grpc.Server {
	unaryInterceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_recovery.UnaryServerInterceptor()))
	streamInterceptors := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_recovery.StreamServerInterceptor()))
	server := grpc.NewServer(unaryInterceptors, streamInterceptors)
	server.RegisterService(&cart.CartSvc_ServiceDesc, impl)
	return server

}

func NewRestGateway(grpcDestPort int, opts ...grpc.DialOption) *http.ServeMux {
	ctx := context.Background()

	grpcMux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(anyHeaderMatcher),
		runtime.WithIncomingHeaderMatcher(anyHeaderMatcher),
	)

	grpcAddress := fmt.Sprintf("localhost:%d", grpcDestPort)
	if err := cart.RegisterCartSvcHandlerFromEndpoint(ctx, grpcMux, grpcAddress, opts); err != nil {
		panic(err)
	}
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		grpcMux.ServeHTTP(writer, request)
	})
	return serverMux

}

func anyHeaderMatcher(header string) (string, bool) {
	if header == "Connection" {
		return header, false
	}
	return header, true
}

func (this *GrpcServer) Start() error {
	grpcLis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", this.grpcPort))
	if err != nil {
		return err
	}

	restLis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", this.restPort))
	if err != nil {
		return err
	}

	ch := make(chan error, 2)

	go func() {
		if err := this.grpcServer.Serve(grpcLis); err != nil {
			ch <- err
		}
	}()

	go func() {
		if err := http.Serve(restLis, this.restServer); err != nil {
			ch <- err
		}
	}()

	for {
		select {
		case err := <-ch:
			log.Println("service crushed")
			return err
		case <-time.After(10 * time.Second):
			log.Println("service heart beat")
			break
		}
	}
}
