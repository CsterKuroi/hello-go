package main

import (
    "fmt"
    "time"
)

var over = make(chan bool)
var sem = make(chan int, 4) //控制并发任务数

func Worker(i int) {
    sem <- 1
    fmt.Println(time.Now().Format("04:05"), i)
    time.Sleep(1 * time.Second)

    <-sem

    if i == 9 {
        over <- true
    }
}

func main() {

    for i := 0; i < 10; i++ {
        go Worker(i)
    }

    <-over

}
