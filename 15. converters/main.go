package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64 = 123

	// Convert int to string
	a := strconv.FormatInt(76, 10)
	a3 := strconv.Itoa(int(n))
	a2 := fmt.Sprint(34)

	fmt.Printf("%s, %T \n", a, a)
	fmt.Printf("%s, %T \n", a2, a2)
	fmt.Printf("%s, %T \n", a3, a3)

	// Convert string to Int
	n1, _ := strconv.ParseInt("225", 10, 64)
	n2, _ := strconv.Atoi("189")

	fmt.Printf("%d, %T \n", n1, n1)
	fmt.Printf("%d, %T \n", n2, n2)

}
