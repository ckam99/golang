package channels

type User struct {
	Name string
	Age  int
}

var UserChannel = make(chan User)
