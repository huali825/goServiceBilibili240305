package main

import "fmt"

func main() {
	server := NewServer("127.0.0.1", 8080)
	server.Start()
	fmt.Println("hello world go")
}
