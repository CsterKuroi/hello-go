package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			continue
			return
		}
	}
}

func main() {
	c := make(chan int, 10)
	quit := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	go func() {
		time.Sleep(time.Microsecond * 1000)
		quit <- 0
	}()
	time.Sleep(time.Second)
	go fibonacci(c, quit)
	time.Sleep(time.Second * 5)
}
