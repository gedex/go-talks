// +build OMIT

package main

import "fmt"

// main START OMIT
func main() {
	// Membentuk dari slice literal
	b := []bool{true, false, true}
	n := []int{1, 2, 3}
	fmt.Println(b, n)

	s := []struct {
		nama string
		umur uint8
	}{
		{"Budi", 17},
		{"Eko", 20},
	}
	fmt.Println(s)
}

// main END OMIT
