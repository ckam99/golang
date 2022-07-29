package main

import (
	"fmt"
	"time"
)

func main() {

	dateString := "2021-11-22"
	dateString = "2022-07-25 08:36:18.802851+00"

	date, error := time.Parse("2006-01-02 15:04:05.999999999MST", dateString)
	//date, error := time.Parse(time.RFC3339, dateString)

	if error != nil {
		fmt.Println("error:", error)
		return
	}

	fmt.Printf("Type of dateString: %T\n", dateString)
	fmt.Printf("Type of date: %T\n", date)
	fmt.Println()
	fmt.Printf("Value of dateString: %v\n", dateString)
	fmt.Printf("Value of date: %v\n", date)
}
