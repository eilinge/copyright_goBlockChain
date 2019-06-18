package main

import (
	"fmt"
	"log"
	"net"
)

var (
	ch_login  chan string
	ch_logout chan string
	ch_msg    chan string
)

var mapClients map[string]net.Conn

func sendMsg(conn net.Conn, msg string) {
	// defer conn.Close()
	num, err := conn.Write([]byte(msg))
	if num <= 0 || err != nil {
		fmt.Println("conn.Write failed")
		delete(mapClients, conn.RemoteAddr().String())
		ch_logout <- conn.RemoteAddr().String() + "conn failed, logout"
		return
	}
}

func Control() {
	// 遍历climaps, 逐一往chan中写入数据
	var msg string
	for {
		select {
		case msg = <-ch_login:
		case msg = <-ch_login:
		case msg = <-ch_msg:
		}

		for _, conn := range mapClients {
			go sendMsg(conn, msg)
		}
	}

}

func commFromClient(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	// 1. 接收消息
	// 2. 将消息写入消息通道
	// 3. 如果有异常, 断开, 清除map
	for {
		n, err := conn.Read(buf) // 读取cli发送来的message
		if err != nil || n <= 0 {
			fmt.Println("conn.Read err: ", err)
			break
		}
		ch_msg <- conn.RemoteAddr().String() + string(buf[:n])
	}
	delete(mapClients, conn.RemoteAddr().String())
	ch_logout <- conn.RemoteAddr().String() + "conn failed, logout"
}

func main() {
	//1. 初始化: 通道和map
	ch_login = make(chan string)
	ch_logout = make(chan string)
	ch_msg = make(chan string)
	mapClients = make(map[string]net.Conn)
	//2. listen 启动监听
	cli, err := net.Listen("tcp", "localhost:8080")
	fmt.Println("the net is listening: localhost:8080")
	if err != nil {
		log.Fatal("net listen err: ", err)
	}
	defer cli.Close()

	go Control()
	//3. 循环等待客户端连接
	for {
		conn, err := cli.Accept()
		if err != nil {
			fmt.Println("net listen err: ", err)
			continue
		}
		// defer conn.Close()
		// 得到新的连接-> 创建通信协程-> 发送到登录所用的通道
		mapClients[conn.RemoteAddr().String()] = conn
		ch_login <- conn.RemoteAddr().String() + "login success"

		go commFromClient(conn)
	}
}
