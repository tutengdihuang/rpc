package client

import (
	server "01ForTest/03RPC/03Asynchronous_call/server"
	"fmt"
	"log"
	"net/rpc"
	"reflect"
)

func ClientAsync() {
	client, err := rpc.DialHTTP("tcp", ":20000")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	argsMul := &server.Args{9, 9}
	var reply int
	err = client.Call("Arith.Multiply", argsMul, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", argsMul.A, argsMul.B, reply)

	//重点讲这里。
	// Asynchronous call
	argsDiv := server.Args{7, 8}
	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", argsDiv, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
	fmt.Println("replayCall")
	fmt.Println("replayCall的数据类型,请看下边type Call struct：", reflect.TypeOf(replyCall))
	fmt.Println("调用的方法：", replyCall.ServiceMethod)
	fmt.Println("返回值：", replyCall.Reply)
	fmt.Println("服务器端完成操作返回：", replyCall.Done)
	fmt.Println("client请求参数：", replyCall.Args)
	fmt.Println("服务器返回的错误：", replyCall.Error)
}

//
/*type Call struct {
	ServiceMethod string      // The name of the service and method to call.
	Args          interface{} // The argument to the function (*struct).
	Reply         interface{} // The reply from the function (*struct).
	Error         error       // After completion, the error status.
	Done          chan *Call  // Strobes when call is complete.
}*/
