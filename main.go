package main

import (
	"fmt"
  "main/internal/domain/books"
  "encoding/json"
  
)

func main() {
  b := books.Book{}
	fmt.Println(b)
  pjson(b)
}

func pjson(a interface{}){
  b,_:= json.MarshalIndent(a,""," ")
  fmt.Println(string(b))
}