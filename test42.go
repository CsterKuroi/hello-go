package main

import (
	"fmt"
	"cryptoconditions"
)

func main() {
        var pub string
        var pri string
	pri, pub = cryptoconditions.GenerateKeypair()
	fmt.Printf("%s,%s\n", pri, pub)
}
