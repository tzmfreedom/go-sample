package main

import (
	"context"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/tzmfreedom/go-sample/protobuf/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	go s.Serve(lis)

	wrappedGrpc := grpcweb.WrapServer(s, grpcweb.WithOriginFunc(func(origin string) bool { return true }))
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		} else if wrappedGrpc.IsAcceptableGrpcCorsRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}
		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(resp, req)
	}))
}

type server struct {}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: in.GetName() + ":reply"}, nil
}
