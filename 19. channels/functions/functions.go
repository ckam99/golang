package functions

import (
	"math/rand"
	"time"
)

func Randomize(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
