package main

import "fmt"

type User struct {
	Name     string
	Username string
}

func main() {
	u := User{
		Name:     "Akeda",
		Username: "gedex",
	}
	fmt.Println(u)
}
