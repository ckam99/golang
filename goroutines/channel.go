package main

import (
	"fmt"
	"time"
)

var ch chan string = make(chan string, 2)
var ch2 chan int = make(chan int, 3)

func Task(delay time.Duration) {
	time.Sleep(delay * time.Second)
	ch <- fmt.Sprintf("%s %s", delay*time.Second, "hello")
	ch2 <- int(delay * time.Second)
}

func TrackChannel() {
	go func() {
		for s := range ch {
			fmt.Println(s)
		}
	}()

	go func() {
		for s := range ch2 {
			fmt.Println(s)
		}
	}()

	// select {
	// case out := <-ch:
	// 	fmt.Println(out)
	// case <-time.After(10 * time.Second):
	// 	fmt.Println("timeout....1")
	// }
}

func ChannelExample() {

	go Task(5)
	go Task(2)

	go TrackChannel()
	fmt.Println("Gorotines")
	time.Sleep(8 * time.Second)

	defer close(ch)

}
