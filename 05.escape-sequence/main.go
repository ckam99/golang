package main

import "fmt"

/*

\\	\ character
\'	' character
\"	" character
\?	? character
\a	Alert or bell
\b	Backspace
\f	Form feed
\n	Newline
\r	Carriage return
\t	Horizontal tab
\v	Vertical tab
\ooo	Octal number of one to three digits
\xhh . . .	Hexadecimal number of one or more digits

*/
func main() {
	fmt.Printf("Hello\aWorld!\n")
	fmt.Printf("Hello\bWorld!\n")
	fmt.Printf("Hello\fWorld!\n")
	fmt.Printf("Hello\nWorld!\n")
	fmt.Printf("Hello\rWorld!\n")
	fmt.Printf("Hello\tWorld!\n")
	fmt.Printf("Hello\vWorld!\n")

}
