package main

import (
"fmt"
"time"
)

func main() {
ch := make(chan int)
go ff(ch)
fmt.Println("1", time.Now())
ch <- 2
fmt.Println("2", time.Now())
fmt.Println("main end")
fmt.Println("5", time.Now())
}

func ff(ch chan int) {
fmt.Println("3", time.Now())
time.Sleep(2 * time.Second)
<-ch
fmt.Println("ff end")
fmt.Println("4", time.Now())
}
