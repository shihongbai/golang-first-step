package main

import (
	"github.com/shihongbai/golang-first-step/about_project_demo/about_web/golang-webservice-example/controller"
	"log"
	"net/http"
)

func main() {
	// 初始化路由
	http.HandleFunc("/sayHello", controller.SayHelloHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/insert_user", controller.InsertUserHandler)

	// 启动服务
	log.Println("Starting server on :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
