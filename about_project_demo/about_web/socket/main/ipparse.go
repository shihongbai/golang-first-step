package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr != nil {
		fmt.Println("IP:", addr.String())
	} else {
		fmt.Println("Invalid IP:", name)
	}

	tcp, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 80})
	if err != nil {
		return
	}

	tcp.Accept()

	os.Exit(0)
}
