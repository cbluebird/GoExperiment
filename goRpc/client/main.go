package main

import (
	"log"
	"net/rpc"
)

type ArgsTwo struct {
	X, Y int
}

func main() {
	clinet, err := rpc.DialHTTP("tcp", "127.0.0.1:8808")
	if err != nil {
		log.Fatal(err)
	}
	i1 := 1
	i2 := 90
	args := ArgsTwo{i1, i2}
	var reply int
	err = clinet.Call("Algorithm.Sum", args, &reply)
	log.Println(reply)
}
