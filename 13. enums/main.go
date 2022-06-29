package main

import "fmt"

type Profession int

const (
	Doctor Profession = iota
	Developer
	Teacher
	Entrepreneur
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	var myProfession Profession = Developer

	if myProfession == Developer {
		fmt.Println("My Profession is Developer", myProfession)
	}

	// const
	fmt.Println(Sunday)    // вывод 0
	fmt.Println(Wednesday) // вывод 3
	fmt.Println(Saturday)  // вывод 6
}
