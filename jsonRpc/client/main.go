package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Send struct {
	X,Y string
}

func main()  {
	client,err:=jsonrpc.Dial("tcp","127.0.0.1:8085")
	if err != nil {
		log.Println(err)
	}
	send:=Send{"hello","world"}
	var reply string
	err=client.Call("Programmer.GetSkill",send,&reply)
	log.Println(reply)
}
