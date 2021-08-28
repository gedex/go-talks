// +build OMIT

package main

import "fmt"

// main START OMIT
func main() {
	i := 42

	p := &i         // p merupakan pointer i
	fmt.Println(p)  // p memegang alamat i
	fmt.Println(*p) // membaca nilai i melalui pointer p
	*p = 21         // merubah nilai i melalui pointer p
	fmt.Println(i)
}

// main END OMIT
