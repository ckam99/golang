package main

import "fmt"

func updateByPointer(v *int) {
	*v = 8
}

func update(v int) {
	v = 5
}

func main() {
	n := 10
	n2 := 20

	update(n)
	fmt.Println(n) // 10  no change

	updateByPointer(&n2)
	fmt.Println(n2) // 8 changed by pointer
}
