package main

import "fmt"

func main() {

	/*
		        ======== General ===================
				%v	the value in a default format when printing structs, the plus flag (%+v) adds field names
				%#v	a Go-syntax representation of the value
				%T	a Go-syntax representation of the type of the value
				%%	a literal percent sign; consumes no value

				====  The default format for %v is: =====

				bool:                    %t
				int, int8 etc.:          %d
				uint, uint8 etc.:        %d, %#x if printed with %#v
				float32, complex64, etc: %g
				string:                  %s
				chan:                    %p
				pointer:                 %p

				======== Boolean ==================
				%t	the word true or false

				======== Integers =================
				%b	base 2
				%c	the character represented by the corresponding Unicode code point
				%d	base 10
				%o	base 8
				%O	base 8 with 0o prefix
				%q	a single-quoted character literal safely escaped with Go syntax.
				%x	base 16, with lower-case letters for a-f
				%X	base 16, with upper-case letters for A-F
				%U	Unicode format: U+1234; same as "U+%04X"

				========= Floating-point and complex constituents ===========
				%b	decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format, e.g. -123456p-78
				%e	scientific notation, e.g. -1.234456e+78
				%E	scientific notation, e.g. -1.234456E+78
				%f	decimal point but no exponent, e.g. 123.456
				%F	synonym for %f
				%g	%e for large exponents, %f otherwise. Precision is discussed below.
				%G	%E for large exponents, %F otherwise
				%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
				%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

				======== String and slice of bytes (treated equivalently with these verbs) =========
				%s	the uninterpreted bytes of the string or slice
				%q	a double-quoted string safely escaped with Go syntax
				%x	base 16, lower-case, two characters per byte
				%X	base 16, upper-case, two characters per byte

				=====  Slice  =============
				%p	address of 0th element in base 16 notation, with leading 0x

				======== Pointer ===========
				%p	base 16 notation, with leading 0x
				The %b, %d, %o, %x and %X verbs also work with pointers,
				formatting the value exactly as if it were an integer.

	*/

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
