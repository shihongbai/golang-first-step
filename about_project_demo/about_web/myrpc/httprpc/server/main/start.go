package main

import (
	"fmt"
	"golang-first-step/about_project_demo/about_web/myrpc/httprpc/server"
	"net/http"
	"net/rpc"
)

// RPC服务端注册
// 我们注册了一个Arith的RPC服务，然后通过rpc.HandleHTTP
// 函数把该服务注册到了HTTP协议上，然后我们就可以利用http的方式来传递数据了。
func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
	}
}
