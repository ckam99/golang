package main

import "fmt"

var x int = 45

func main() {
	PrintX()
	x = 23
	Test()
	PrintX()
	F()
	F2()
	PrintX()
}

func PrintX() {
	fmt.Printf("global x %v \n", x)
}

func F() {
	x = 89
	fmt.Printf("global x from F %v \n", x)
}

func F2() {
	x := 54
	fmt.Printf(" x from F2 %v \n", x)
}
