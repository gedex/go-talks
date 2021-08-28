package main

import "fmt"

// main START OMIT
var m map[string]int

func main() {
	m := map[string]int{
		"Januari":  1,
		"Februari": 2,
	}

	fmt.Println(m)

	delete(m, "Januari")

	v, ok := m["Januari"]
	fmt.Println("v: ", v, "ok: ", ok)
}

// main END OMIT
