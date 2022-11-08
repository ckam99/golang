package main

import (
	"github.com/ckam225/go-collection"
)

func main() {
	numbers := []int{3, 6, 8, 9}

	collection.Collect(numbers).Filter(func(currentNumber int, index int64) bool {
		return currentNumber > 7
	}).ToList()
	// output: 8,9

	collection.Collect(numbers).Map(func(currentNumber int, index int64) int {
		return currentNumber * 2
	}).ToList()
	// output: 6,12,16,18

	collection.Collect(numbers).Remove(2).ToList()
	// output: 3,6,9

	collection.Collect(numbers).Shift().ToList()
	// output: 6,8,9

	collection.Collect(numbers).Pop().ToList()
	// output: 3,6,8

	collection.Collect(numbers).Join("-")
	// output: 3-6-8-9
}
