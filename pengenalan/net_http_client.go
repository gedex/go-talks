package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// main START OMIT
func main() {
	res, err := http.Get("https://covid19.mathdro.id/api/countries/id")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", body)
}

// main END OMIT
