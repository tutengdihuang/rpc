package main

import (
	"fmt"
	_ "net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Args struct {
	A, B int
}

const (
	URL = "127.0.0.1:15001"
)

func main() {

	client, err := jsonrpc.Dial("tcp", URL)
	defer client.Close()

	if err != nil {
		fmt.Println(err)
	}

	var args [100000]Args //批量参数
	for i := 0; i < len(args); i++ {
		args[i].A = i
		args[i].B = i
	}

	//fmt.Printf("args %v \n", args)

	harvester_time1 := time.Now().UnixNano()
	var reply int

	err = client.Call("Arith.Add", args, &reply) //批量执行
	if err != nil {
		fmt.Println(err)
	}
	harvester_time2 := time.Now().UnixNano()

	fmt.Printf("cal time %f \n", float64(harvester_time2-harvester_time1)/float64(1000000000))
	fmt.Println(reply)
}
