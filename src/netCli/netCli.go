package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("net.Dial failed")
	}
	defer conn.Close()
	go func() {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) // 读取server发送来的message
		if err != nil {
			fmt.Println("the buf read err", err)
		}
		fmt.Println("receive message is: ", string(buf[:n]))
	}()
	// conn.Write([]byte("eilinge"))
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf) // 读取到键盘输入
		if err != nil {
			fmt.Println("net.Dial failed")
			break
		}
		conn.Write(buf[:n]) // 响应给server的message
	}
}
