package main

import "fmt"

func main() {
	var (
		a  bool
		b  string
		c  int
		d  uint
		Pi float32
		e  complex64
	)
	fmt.Println("Data types")

	/*
					   ======================
					   === Integer Types ====
					   ======================

				       uint8 | byte: Unsigned 8-bit integers (0 to 255)
					   uint16:  Unsigned 16-bit integers (0 to 65535)
					   uint32: Unsigned 32-bit integers (0 to 4294967295)
					   uint64: Unsigned 64-bit integers (0 to 18446744073709551615)
					   int8: Signed 8-bit integers (-128 to 127)
					   int16: Signed 16-bit integers (-32768 to 32767)
					   int32 | rune: Signed 32-bit integers (-2147483648 to 2147483647)
					   int64: Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
					   uint : 32 or 64 bits
					   int: same size as uint
					   uintptr: an unsigned integer to store the uninterpreted bits of a pointer value

					   =======================
					   === Floating Types ====
					   =======================
					   float32:  IEEE-754 32-bit floating-point numbers
		               float64:  IEEE-754 64-bit floating-point numbers
		               complex64:  Complex numbers with float32 real and imaginary parts
		               complex128:  Complex numbers with float64 real and imaginary parts
	*/

	a = true
	b = "Hello"
	c = -79
	d = 45
	e = 2i - 4
	Pi = 3.14

	fmt.Printf("a is %T \n", a)
	fmt.Printf("b is %T \n", b)
	fmt.Printf("c is %T \n", c)
	fmt.Printf("d is %T \n", d)
	fmt.Printf("e is %T \n", e)
	fmt.Printf("PI is %T \n", Pi)

}
