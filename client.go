package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	//建立网络链接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络连接失败")
	}

	var pd int
	/*
		func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {

	*/
	err = cli.Call("Panda.Getinfo", 10086, &pd)
	if err != nil {
		fmt.Println("打call失败")
	}

	fmt.Println("最后得到的值为：", pd)

}
