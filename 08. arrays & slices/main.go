package main

import (
	"fmt"
	"strings"
)

func main() {

	var arr [5]int
	arr[0] = 3
	arr[1] = 8
	arr[2] = 0
	arr[3] = 13
	arr[4] = 10

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	list := [...]int{9, 8, 4, 9, 0}
	otherList := [2]string{"Rachelle", "Dialo"}

	for k, v := range list {
		fmt.Printf("Key=%v , value=%v \n", k, v)
	}

	for _, value := range otherList {
		fmt.Printf("%v \n", value)
	}

	for _, v := range strings.Split("abacef", "") {
		fmt.Printf("%v \n", v)
	}

	// differences

	fixedArray := [...]int{2, 4, 6, 8, 10}
	fmt.Println(fixedArray)
	// fixedArray[5] = 12 // dontdo that
	// append(fixedArray, 12)

	// Slices
	sliceArray := []int{1, 3, 5, 7, 9}
	fmt.Println(sliceArray)
	// sliceArray[5] = 11 // be careful
	sliceArray = append(sliceArray, 11, 13, 15)
	fmt.Println(sliceArray)

	otherSlice := make([]int, 8)
	fmt.Println(otherSlice)

	fmt.Println(sliceArray[2:5])
	fmt.Println(fixedArray[2:])

	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Базовый массив: %v\n", baseArray)

	baseSlice := baseArray[5:8]
	fmt.Printf(
		"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v \n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)
}
