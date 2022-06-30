package main

import "fmt"

type Animal struct {
	name   string
	family string
}

func (a Animal) Describe() {
	fmt.Println(a.name, "is", a.family)
}

func main() {
	lion := Animal{name: "Lion", family: "Felin"}
	lion.Describe()
}
