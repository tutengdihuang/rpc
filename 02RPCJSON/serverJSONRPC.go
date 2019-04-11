package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const (
	URL = "127.0.0.1:15001"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Add(args []Args, reply *int) error {
	for i := 0; i < len(args); i++ {
		*reply += args[i].A
		*reply += args[i].B
	}
	return nil
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", URL)
	fmt.Println(URL)
	if err != nil {
		fmt.Println(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
