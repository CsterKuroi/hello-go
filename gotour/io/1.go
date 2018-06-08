package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("tmp/dat") //ioutil读文件
	check(err)
	fmt.Println(string(dat))
	f, err := os.Open("tmp/dat") //os file打开文件
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // file reader
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
	o2, err := f.Seek(7, 0) //seek
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	o3, err := f.Seek(7, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f) //bufio reader
	b4, err := r4.Peek(5) //peek
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()
}
