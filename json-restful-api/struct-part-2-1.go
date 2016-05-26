package main

import "fmt"

type User struct {
	Name     string
	Username string
}

// START OMIT
func main() {
	u := NewUser("Akeda Bagus", "gedex")
	fmt.Println(u)
}

func NewUser(name, username string) *User {
	user := new(User)
	user.Name = name
	user.Username = username

	return user
}

// END OMIT
