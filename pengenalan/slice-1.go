// +build OMIT

package main

import (
	"fmt"
)

// main START OMIT
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Membentuk slice dari array
	var s []int = primes[1:4]
	fmt.Println(s)

	// Mengubah slice s ikut mengubah array primes
	s[0] = 4
	fmt.Println(s)
	fmt.Println(primes)
}

// main END OMIT
