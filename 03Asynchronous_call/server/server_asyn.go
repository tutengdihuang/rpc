package server

import (
	"log"
	"net"
	"net/http"
	//"net/http"
	"net/rpc"
)

func ServerAsyn() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP() //注册默认的方法。
	l, e := net.Listen("tcp", ":20000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	//关于http server 请查看 https://www.jianshu.com/p/a690cbc67ab7

}
