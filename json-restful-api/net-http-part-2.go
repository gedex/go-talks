package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/articles/", articlesHandler)
	mux.HandleFunc("/users/", usersHandler)
	http.ListenAndServe(":8080", mux)
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "articles")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "users")
}

// END OMIT
