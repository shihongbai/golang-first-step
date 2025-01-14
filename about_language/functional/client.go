package main

import (
	"golang-first-step/about_language/functional/model"
	"golang-first-step/about_language/functional/option"
)

// functional option编程模式的好处
/**
1. 直觉式编程
2. 高度的可配置化
*/
func main() {
	server := model.NewServer("localhost", 8080, option.MaxConns(1000), option.Timeout(100))

	println(server.ToString())
}
