package main

import (
	"fmt"
	"time"
)

// START OMIT
type User struct {
	Name     string
	Username string
}

type Article struct {
	Title  string
	Body   string
	Date   time.Time
	Author User
}

func main() {
	u := User{"Akeda Bagus", "gedex"}
	a := new(Article)
	a.Author = u

	fmt.Println(a)
}

// END OMIT
