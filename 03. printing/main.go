package main

import "fmt"

func main() {

	var (
		a     bool      = false
		b     string    = "hello"
		c     int       = -30045
		d     uint      = 255
		e     uint8     = 255
		f     float64   = 56.485408058344804
		Pi    float32   = 3.14
		compl complex64 = 2i - 1
	)

	fmt.Printf("a=%t is %T \n", a, a)
	fmt.Printf("b=%s is %T \n", b, b)
	fmt.Printf("c=%d is %T \n", c, c)
	fmt.Printf("d=%d is %T \n", d, d)
	fmt.Printf("e=%d is %T \n", e, e)
	fmt.Printf("f=%f is %T \n", f, f)
	fmt.Printf("f=%0.2f is %T \n", f, f)
	fmt.Printf("Pi=%g is %T \n", Pi, Pi)
	fmt.Printf("compl=%g is %T \n", compl, compl)
}
