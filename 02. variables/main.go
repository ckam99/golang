package main

import "fmt"

func main() {

	const company string = "Microsoft"
	var firstname, lastname string

	var (
		gender string
		age    int
	)
	salary := 9000.45

	firstname, lastname = "Claver", "Amon"
	gender = "male"
	age = 29

	fmt.Printf("My name is %s %s, i'm %d years old, %s. I am working in %s and i win %f per month\n", firstname, lastname, age, gender, company, salary)

}
