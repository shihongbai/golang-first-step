package controller

import (
	"database/sql"
	"fmt"
	"golang-first-step/about_project_demo/about_web/golang-webservice-example/config"
	"net/http"
)

var db sql.DB // var to connect with database

func init() {
	db = config.Connect() // connect DB while server is On
}

// User 用户结构体
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// SayHelloHandler 处理 /sayHello 接口
func SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// LoginHandler 处理 /login 接口
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 如果是 POST 请求，则处理登录表单提交
	if r.Method == http.MethodPost {
		// 获取表单提交的用户名和密码
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 根据用户名从数据库中查询用户信息
		var dbUser User
		err := db.QueryRow("SELECT id, username, password FROM `user` WHERE username=?", username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
		if err != nil {
			// 用户不存在或者查询出错，返回登录失败
			fmt.Fprintf(w, "Login failed: %v", err)
			return
		}

		// 校验密码是否匹配
		if password != dbUser.Password {
			// 密码不匹配，返回登录失败
			fmt.Fprintf(w, "Login failed: Invalid password")
			return
		}

		// 登录成功
		fmt.Fprintf(w, "Login successful! Welcome, %s", dbUser.Username)
		return
	}

	// 如果是 GET 请求，则返回登录表单页面
	// 注意：这里只是简单的示例，实际中需要一个美观的 HTML 登录页面
	fmt.Fprintf(w, `
        <html>
            <head>
                <title>Login</title>
            </head>
            <body>
                <h2>Login</h2>
                <form method="post" action="/login">
                    <label for="username">Username:</label><br>
                    <input type="text" id="username" name="username"><br>
                    <label for="password">Password:</label><br>
                    <input type="password" id="password" name="password"><br><br>
                    <input type="submit" value="Login">
                </form>
            </body>
        </html>
    `)
}

// InsertUserHandler 处理 /insert_user 接口
func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	// 如果是 POST 请求，则处理新增用户表单提交
	if r.Method == http.MethodPost {
		// 获取表单提交的用户名和密码
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 插入新用户到数据库
		_, err := db.Exec("INSERT INTO user(username, password) VALUES(?, ?)", username, password)
		if err != nil {
			// 插入失败，返回错误信息
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 插入成功，返回成功状态
		fmt.Fprintf(w, "User inserted successfully")
		return
	}

	// 如果是 GET 请求，则返回新增用户表单页面
	// 注意：这里只是简单的示例，实际中需要一个美观的 HTML 表单页面
	fmt.Fprintf(w, `
        <html>
            <head>
                <title>Insert User</title>
            </head>
            <body>
                <h2>Insert User</h2>
                <form method="post" action="/insert_user">
                    <label for="username">Username:</label><br>
                    <input type="text" id="username" name="username"><br>
                    <label for="password">Password:</label><br>
                    <input type="password" id="password" name="password"><br><br>
                    <input type="submit" value="Insert">
                </form>
            </body>
        </html>
    `)
}
