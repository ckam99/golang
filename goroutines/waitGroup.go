package main

import (
	"fmt"
	"sync"
	"time"
)

func WriteSomething(title string, wg *sync.WaitGroup) {
	for i := 0; i < 9; i++ {
		fmt.Println(title)
		time.Sleep(500 * time.Millisecond)
	}
	defer wg.Done()
}

func WaitGroupExample() {
	var wg sync.WaitGroup
	wg.Add(2)
	go WriteSomething("first", &wg)
	go WriteSomething("second", &wg)
	wg.Wait()
}
