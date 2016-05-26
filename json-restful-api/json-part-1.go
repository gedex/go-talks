package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string
	Username string
}

// START OMIT
func main() {
	u := User{
		Name:     "Akeda Bagus",
		Username: "gedex",
	}

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	// First, create placeholder where decoded data will be stored.
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

// END OMIT
