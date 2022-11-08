package main

import (
	"fmt"

	"github.com/ckam225/go-collection"
)

func main() {
	numbers := []int{3, 6, 8, 9}
	fmt.Println(numbers)
	arr := collection.Collect(numbers).Remove(2).ToList()
	fmt.Println(arr)
}
