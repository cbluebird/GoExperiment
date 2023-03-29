package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ArgsList struct {
	X, Y string
}

type Programmer string

func (m *Programmer) GetSkill(a1 *ArgsList, reply *string) error {
	*reply = a1.X + a1.Y
	return nil
}
func main() {
	str := new(Programmer)
	rpc.Register(str)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8085")
	if err != nil {
		log.Println(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}
	jsonrpc.ServeConn(conn)
}
