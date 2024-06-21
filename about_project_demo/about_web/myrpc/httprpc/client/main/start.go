package main

import (
	"fmt"
	"golang-first-step/about_project_demo/about_web/myrpc/httprpc/client"
	"log"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	cli, err := rpc.DialHTTP("tcp", serverAddress+":12345")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 客户端rpc调用
	args := client.Args{17, 8}
	var reply int
	err = cli.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot client.Quotient
	err = cli.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d\n", args.A, args.B, quot)

}
