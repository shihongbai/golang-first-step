package main

import (
	myGrpc "golang-first-step/about_project_demo/about_rpc/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 服务端启动入口
func main() {
	grpcServer := grpc.NewServer()
	myGrpc.RegisterHelloServiceServer(grpcServer, new(myGrpc.HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
