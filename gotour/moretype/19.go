package main

import "fmt"

func adder() (mm func(x int) int) {
	sum := 0
	mm = func(x int) int {
		sum += x
		return sum
	}
	return
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
