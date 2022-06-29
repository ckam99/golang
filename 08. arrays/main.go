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

}
