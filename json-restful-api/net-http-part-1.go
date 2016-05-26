package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello GoJakarta!")
}
