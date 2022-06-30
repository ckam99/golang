package main

import (
	"fmt"
	f "gochanels/functions"
)

func Calc(myChannel chan int) {
	rn := f.Randomize(90)
	myChannel <- rn
}

func main() {
	foo := make(chan int)
	defer close(foo) // not recommanded in main function
	go Calc(foo)
	number := <-foo
	fmt.Println(number)
}
