package main

import "fmt"

// main START OMIT
var m map[string]int

func main() {
	// Inisialisasi dengan make
	m = make(map[string]int)
	m["satu"] = 1
	m["dua"] = 2
	fmt.Println(m)

	// Dengan literal map
	m2 := map[string]int{
		"Januari":  1,
		"Februari": 2,
	}
	fmt.Println(m2)
}

// main END OMIT
