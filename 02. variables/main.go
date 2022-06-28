package main

import "fmt"

func main() {

	var firstname, lastname string

	var (
		gender string
		age    int
	)
	salary := 9000.45

	firstname = "Claver"
	lastname = "Amon"
	gender = "male"
	age = 29

	fmt.Printf("My name is %s %s, i'm %d years old, %s. I win %f per month\n", firstname, lastname, age, gender, salary)
}
