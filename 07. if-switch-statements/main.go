package main

import (
	"fmt"
	"strings"
)

const ci string = "Ghana"
const ru string = "Russia"
const ca string = "Canada"

func main() {
	var age int = 23

	fmt.Print("How year old ? ")
	fmt.Scanf("%d", &age)

	if age < 18 {
		fmt.Println("You are mom")
	} else if age >= 18 && age <= 60 {
		fmt.Println("You are young")
	} else {
		fmt.Println("You are adult")
	}

	var country string
	fmt.Print("Enter your country: ")
	fmt.Scanf("%s", &country)

	println(country)
	switch strings.ToLower(country) {
	case strings.ToLower(ci):
		fmt.Println("You are ivorian")
	case strings.ToLower(ca):
		fmt.Println("You are canadian")
	case strings.ToLower(ru):
		fmt.Println("You are russian")
	default:
		fmt.Println("Option not available")
	}

}
