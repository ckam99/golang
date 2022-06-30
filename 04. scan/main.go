package main

import "fmt"

func main() {
	var lastname, firstname string
	var age int

	fmt.Println("Enter your lastname and firstname: ")
	fmt.Scan(&lastname, &firstname)

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Hi", firstname, lastname, "you are", age, "years old")
}
