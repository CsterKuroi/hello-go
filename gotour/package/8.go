package main

import "fmt"

var i, j int = 1, 2
var (
	a        = 1
	b uint64 = 3
	e        = 3.14
	d        = 4
)

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
