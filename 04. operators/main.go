package main

import "fmt"

func main() {

	a, b, c := 10, 2, 10

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	fmt.Printf("%v / %v = %v\n", a, b, a/b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	fmt.Printf("%v %% %v = %v\n", a, b, a%b)

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	fmt.Println(a <= b || a == c)
	fmt.Println(a >= b && a == b)
}
