package main

import "fmt"

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Can not divide number by zero")
	}
	return x / y
}

func safeExit() {
	if r := recover(); r != nil {
		fmt.Printf("Panic is recovered\n")
	}
}

func main() {
	defer safeExit()

	fmt.Println(divide(3, 2))
	fmt.Println(divide(3, 0))
	fmt.Println(divide(0, 12))

}
