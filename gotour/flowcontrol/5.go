//package main
//
//import "fmt"
//
//func main() {
//	defer fmt.Println("world")
//
//	fmt.Println("hello")
//}

package main

import "fmt"


func add() int{
	defer fmt.Println("down")
	return 1+1
}

func main() {
	fmt.Println(add())
}
