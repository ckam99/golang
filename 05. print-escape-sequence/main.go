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

	fmt.Println("============================")
	fmt.Print("hello, world")
	fmt.Print("hello, world")
	fmt.Println("hello, world")
	fmt.Print("hello, world")

	s := "hello world"
	fmt.Println(s[:])
	fmt.Println(s[1:])
	fmt.Println(s[1:4])

	abc := "abc"

	b := []byte(abc)
	for _, k := range b {
		fmt.Print(string(k))
		fmt.Printf(" %c\n", k)
	}

	fmt.Printf("\\")
	fmt.Printf("\"")
	fmt.Printf("123\f456\f789")
	fmt.Printf("\vf\v")

	fmt.Print("33", 27)   // Ivan27
	fmt.Println("33", 27) // Ivan 27
	fmt.Print(33, 27)     // 33 27
	fmt.Print("33", "27") // 3327
}
