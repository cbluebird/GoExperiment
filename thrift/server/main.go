package main

import (
	"context"
	"fmt"

	"experiment/thrift/gen-go/hello"
	"github.com/apache/thrift/lib/go/thrift"
)

type HelloHandler struct{}

func (h *HelloHandler) Greet(ctx context.Context, req *hello.HelloReq) (*hello.HelloRes, error) {
	fmt.Printf("Received request from %s\n", req.GetName())

	res := &hello.HelloRes{
		Greeting: fmt.Sprintf("Hello, %s!", req.GetName()),
	}

	return res, nil
}

func main() {
	handler := &HelloHandler{}
	processor := hello.NewHelloProcessor(handler)

	serverTransport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Starting the server...")

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
