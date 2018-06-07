package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//ResolveTCPAddr用于获取一个TCPAddr
	//net参数是"tcp4"、"tcp6"、"tcp"
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "www.baidu.com:80")
	chkError(err)

	//DialTCP建立一个TCP连接
	//laddr表示本机地址，一般设为nil  //raddr表示远程地址
	tcpconn, err2 := net.DialTCP("tcp", nil, tcpaddr)
	chkError(err2)

	//向tcpconn中写入数据
	_, err3 := tcpconn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	chkError(err3)

	//读取tcpconn中的所有数据
	data, err4 := ioutil.ReadAll(tcpconn)
	chkError(err4)

	fmt.Println(string(data))
}
