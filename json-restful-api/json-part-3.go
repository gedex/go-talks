package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// START OMIT
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func main() {
	u := User{"Akeda Bagus", "gedex"}
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// END OMIT
