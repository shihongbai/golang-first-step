package golang_IM_system

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}

// 创建一个user
func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr()

	user := &User{
		Name: addr.String(),
		Addr: addr.String(),
		C:    make(chan string),
		conn: conn,
	}
	return user
}
