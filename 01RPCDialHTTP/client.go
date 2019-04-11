package _1RPCDialHTTP

import (
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	//除非设置了显式的编解码器，本包默认使用encoding/gob包来传输数据。
	for {
		time.Sleep(time.Second)
		//建立网络链接
		cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
		if err != nil {
			fmt.Println("网络连接失败")
		}

		var pd int
		/*
			func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
		*/
		err = cli.Call("Allen.Getinfo", 100, &pd)
		if err != nil {
			fmt.Println("打call失败")
		}

		fmt.Println("最后得到的值为：", pd)
	}
}
