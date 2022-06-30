package main

import (
	"fmt"
)

func main() {
	// classical loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Printf("\n\n")

	// while alternative
	x := 0
	for x < 10 {
		fmt.Printf("%v ", x)
		x++
	}
	fmt.Printf("\n\n")

	// break operator
	y := 0
	for {
		if y >= 10 {
			break
		}
		fmt.Printf("%v ", y)
		y++
	}
	fmt.Printf("\n\n")

	// continue operator
	k := 0
	for ; k < 20; k++ {
		if k%2 != 0 {
			continue
		}
		fmt.Printf("%v ", k)
	}
	fmt.Printf("\n\n")
}
