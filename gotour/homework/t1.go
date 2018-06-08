package main

import "fmt"

func main() {
	t1 := []int{1, 2}
	t2 := make([]int, 1, 1)
	t2 = append(t2, t1...)
	fmt.Println(t2)
}
