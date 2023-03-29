package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	X,Y int
}

type Algorithm int

func  (t *Algorithm)Sum(args *Args,reply *int)error{
	*reply=args.X+args.Y
	fmt.Println("sum in server",reply)
	return nil
}

func main()  {
	algorithm:=new(Algorithm)
	fmt.Println("Algorithm start",algorithm)
	rpc.Register(algorithm)
	rpc.HandleHTTP()
	err:=http.ListenAndServe(":8808",nil)
	if err!=nil{
		log.Println(err)
	}
}