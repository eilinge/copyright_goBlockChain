package main

import (
	"fmt"
	"log"
	"net"
)

func main10() {
	cli, err := net.Listen("tcp", "localhost:8080")
	fmt.Println("the net is listening: localhost:8080")
	if err != nil {
		log.Fatal("net listen err: ", err)
	}
	defer cli.Close()

	conn, err := cli.Accept()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect successfully")

	defer conn.Close()
	if err != nil {
		log.Fatal("cli.Accept err: ", err)
	}
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) // 读取cli发送来的message
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			break
		}
		if n > 0 {
			conn.Write(buf[:n]) // 响应给cli的message
		}
		fmt.Println("message: ", string(buf[:n-1])) // 显示在本server/client终端中
	}
}
