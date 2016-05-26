package main

import "fmt"

type User struct {
	Username string
}

func (u *User) Login() {
	fmt.Printf("Logging in %s...", u.Username)
}

type Article struct {
	Title string
}

type Comment struct {
	Body string
}

type Author struct {
	User
	Articles []*Article
	Comments []*Comment
}

// START OMIT
func main() {
	author := new(Author)
	author.Username = "gedex"
	author.Login()
}

// END OMIT
