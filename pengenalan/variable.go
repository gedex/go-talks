// +build OMIT

package main

import "fmt"

var (
	i int
	s string
)

func main() {
	fmt.Println(i, s)
	i = 1
	s = "foo"
	fmt.Println(i, s)

	// Dengan nilai inisial
	var j int = 2
	fmt.Println(j)

	// Deklarasi singkat
	k := 3
	fmt.Println(k)
}
