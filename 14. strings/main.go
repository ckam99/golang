package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hello world"
	// const
	fmt.Println(strings.Contains(str, "wo"))
	fmt.Println(strings.Contains(str, "wire"))
	fmt.Println(strings.ContainsRune(str, 'w'))
	fmt.Println(strings.ContainsRune(str, 'z'))

	fmt.Println(strings.Compare(str, "awerr"))
	fmt.Println(strings.Compare(str, str))

	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToTitle(str))

	fmt.Println(strings.Count(str, "l"))
	fmt.Println(strings.LastIndex(str, "l"))
	fmt.Println(strings.Join(strings.Split(str, ""), "-"))

	fmt.Println(strings.HasPrefix(str, "Hello"))
	fmt.Println(strings.HasPrefix(str, "world"))
	fmt.Println(strings.HasSuffix(str, "world"))

	fmt.Println(strings.Replace(str, "Hello", "hi", 1))
	fmt.Println(strings.Replace(str, "l", "x", 2))
	fmt.Println(strings.ReplaceAll(str, "l", "x"))

	// String builder

	var sb strings.Builder
	fmt.Println("String builder")
	sb.WriteString("ab")

	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println(sb.Write([]byte{34, 67, 90}))
	fmt.Println(sb.String())
	sb.Reset()
}
