package main

import (
	"bufio"
	"fmt"
	"net"


)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		//
		fmt.Println(err)
	}
	fmt.Fprintf(conn, "POST / HTTP/1.0\r\n\r\n")

	status, err := bufio.NewReader(conn).ReadString(']')
	fmt.Println(status,err)
}
