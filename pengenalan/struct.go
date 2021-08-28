// +build OMIT

package main

import "fmt"

// main START OMIT
type Point struct {
	X int
	Y int
}

func main() {
	v := Point{3, 4}
	fmt.Println(v)
	v.Y = 5
	fmt.Println(v)

	// Pointer ke struct
	p := &v
	p.X = 1
	fmt.Println(v)
}

// main END OMIT
