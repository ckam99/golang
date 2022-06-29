package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("\b Can not divide number by zero")
	}
	return a / b, nil
}

func SayHello() {
	fmt.Println("Hello world")
}

func Welcome(name string, age int) {
	fmt.Printf("Hi %v you are %v years old", name, age)
}

func main() {
	SayHello()

	fmt.Println(Divide(3, 9))
	fmt.Println(Divide(5, 0))

	fmt.Println(Add(4, 5))
}
