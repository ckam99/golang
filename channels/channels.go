package channels

import (
	"fmt"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var mutex sync.Mutex
var UserChannel = make(chan User)

func ProduceUser(user User) {
	mutex.Lock()
	UserChannel <- user
	mutex.Unlock()
}

func RegisterConsumer() {
	go func() {

		for i := 0; i < 1; i++ {
			select {
			case user := <-UserChannel:
				fmt.Println(user)
			default:
				fmt.Println("UNKNOW")
			}
		}

	}()
}

func UnregisterConsumer() {
	close(UserChannel)
}
