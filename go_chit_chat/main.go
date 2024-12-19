package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func createPost() {
	reqBody, _ := json.Marshal(map[string]string{"key1": "val1", "key2": "val2"})

	resp, _ := http.Post(":8091", "application/json", bytes.NewReader(reqBody))
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("resp: %s", respBody)
}
