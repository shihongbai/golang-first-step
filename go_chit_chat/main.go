package main

import "net/http"

// golang Web项目：chit chat
func main() {
	//  HTTP 请求多路复用器
	mux := http.NewServeMux()
	// 静态服务文件
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
