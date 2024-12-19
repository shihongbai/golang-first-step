package main

import (
	imServer "golang-first-step/about_project_demo/golang-IM-system"
)

func main() {
	s := imServer.NewServer("127.0.0.1", 8080)

	s.Start()
}
