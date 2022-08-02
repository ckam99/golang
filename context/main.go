package main

import (
	"fmt"
	"os"
)

func main() {
	basedir, _ := os.Getwd()
	fmt.Println("Context", basedir)
}
