package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello word  hello panda")
}

/*
方法是导出的
方法有两个参数，都是导出类型或内建类型
方法的第二个参数是指针
方法只有一个error接口类型的返回值
func (t *T) MethodName(argType T1, replyType *T2) error
*/

type Panda int

//函数关键字（对象）函数名（对端发送过来的内容 ， 返回给对端的内容） 错误返回值
func (this *Panda) Getinfo(argType int, replyType *int) error {

	fmt.Println("打印对端发送过来的内容为：", argType)

	//修改内容值
	*replyType = argType + 12306

	return nil
}

func main() {

	//页面的请求
	http.HandleFunc("/panda", pandatext)
	//将类实例化为对象
	pd := new(Panda)
	//服务端注册一个对象
	rpc.Register(pd)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	http.Serve(ln, nil)
}
