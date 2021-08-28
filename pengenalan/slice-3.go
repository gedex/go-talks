// +build OMIT

package main

import "fmt"

// main START OMIT
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice dari s dengan length 0
	s = s[:0]
	printSlice(s)

	// Menambah length dari slice
	s = s[:4]
	printSlice(s)

	// Buang dua elemen awal
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// main END OMIT
