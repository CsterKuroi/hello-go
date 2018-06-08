package main

import (
	"fmt"
	"time"
)

type jihao interface {
	Age() string
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Age() string {
	return "123"
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
	//ee.Age()
}

func keke(err error) {
	fmt.Println(err.(*MyError).When)
}

func main() {
	if err := run(); err != nil {
		keke(err)
	}
}
