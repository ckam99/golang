package main

import "fmt"

// from go:1.18

/*
 - type parameter(with constraint)
 - type inference
 - type set
*/
func Add[T float64 | int](x, y T) T {
	return x + y
}

func Multiply[T ~float64](x, y T) T {
	return x * y
}

func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func main() {
	type f float64
	var fa f = 4.5

	fmt.Println(Add(5, 8))
	fmt.Println(Add(5.0, 8.6))
	fmt.Println(Multiply(fa, 8.6))

	fmt.Printf("%T, %v\n", Min(3.4, 8.6), Min(3.4, 8.6))
	fmt.Printf("%T, %v\n", Min(3, 8), Min(3, 8))
}
