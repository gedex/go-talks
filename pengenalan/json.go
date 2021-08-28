package main

import (
	"encoding/json"
	"fmt"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// main START OMIT
func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}

// main END OMIT
