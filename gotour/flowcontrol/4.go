package main

import (
	"fmt"
)

func main() {
	fmt.Print("Go runs on ")
	age := 40
	switch {
	case age > 10:
		fmt.Println("OS X.")
	case age > 20:
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", age)
	}
}
