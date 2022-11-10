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
			fmt.Println("ch", s)
		}
	}()

	go func() {
		for s := range ch2 {
			fmt.Println("ch2", s)
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

	for i := 0; i < 5; i++ {
		go Task((time.Duration)(i + 1))
	}

	go TrackChannel()
	fmt.Println("Gorotines")
	time.Sleep(8 * time.Second)

	defer close(ch)

}
