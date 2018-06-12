package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func chkError3(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//单独处理客户端的请求
func clientHandle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new connection")
	conn.Write([]byte("hello " + time.Now().String()))
}

func main() {
	//创建一个TCP服务端
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	chkError3(err)
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError3(err2)
	fmt.Println("listening...")
	//goroutine模拟一个请求

	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}
		//通过goroutine来处理用户的请求
		go clientHandle(conn)
	}
}
