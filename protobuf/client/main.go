package main

import (
	"context"
	"fmt"
	"github.com/tzmfreedom/go-sample/protobuf/helloworld"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	req := helloworld.HelloRequest{Name: "hogehoge"}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}
