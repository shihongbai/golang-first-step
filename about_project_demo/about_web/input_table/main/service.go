package main

import (
	"fmt"
	template "html/template"
	"log"
	"net/http"
	"strings"
)

// http包详细介绍文档：https://juejin.cn/post/7127535913483108360

/*
*
http包的执行流程
1. 端口监听：func ListenAndServe(addr string, handler Handler) error
2. 请求解析：绑定并listen
3. 路由分配：请求并创建连接
4. 响应处理
*/
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello shbai!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		files, _ := template.ParseFiles("about_project_demo/about_web/input_table/login.html")
		files.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if len(username) == 0 || len(password) == 0 {
			fmt.Fprintf(w, "username or password is empty")
		} else {
			fmt.Fprintf(w, "登录成功")
		}
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
