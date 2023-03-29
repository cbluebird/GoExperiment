package main

import (
	"context"
	"fmt"
	"net"

	"experiment/thrift/gen-go/hello"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	transport, err := thrift.NewTSocket(net.JoinHostPort("localhost", "9090"))
	if err != nil {
		panic(err)
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		panic(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := hello.NewHelloClientFactory(transport, protocolFactory)

	req := &hello.HelloReq{
		Name: "world",
	}

	res, err := client.Greet(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response from server: %s\n", res.GetGreeting())
}
