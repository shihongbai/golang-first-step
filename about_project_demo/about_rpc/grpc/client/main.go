package main

import (
	"context"
	"fmt"
	myGrpc "golang-first-step/about_project_demo/about_rpc/grpc"
	"google.golang.org/grpc"
	"log"
)

// 客户端启动
func main() {
	conn, err := grpc.NewClient("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := myGrpc.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &myGrpc.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
