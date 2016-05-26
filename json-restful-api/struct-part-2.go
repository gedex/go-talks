package main

import "fmt"

type User struct {
	Name     string
	Username string
}

func main() {
	// Allocates zeroed storage and returns its address.
	u1 := new(User)

	// Omit field names if you know the order.
	u2 := User{"Akeda Bagus", "gedex"}  // type User
	u3 := &User{"Akeda Bagus", "gedex"} // type *User

	fmt.Println(u1, u2, u3)
}
