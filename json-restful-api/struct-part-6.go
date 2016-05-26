package main

import "fmt"

type User struct {
	Username string
}

func (u *User) Login() {
	fmt.Printf("Logging in %s...", u.Username)
}

type Author struct {
	User
	Articles []*Article
	Comments []*Comment
}
