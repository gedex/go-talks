package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/articles/", articlesHandler)
	mux.HandleFunc("/users/", usersHandler)
	http.ListenAndServe(":8080", mux)
}

// START OMIT
func articlesHandler(w http.ResponseWriter, r *http.Request) {
	ep := strings.TrimPrefix(r.URL.Path, "/articles")
	switch {
	case ep == "/":
		fmt.Fprintln(w, "return list of articles in JSON")
	case strings.HasSuffix(ep, ".json"):
		fmt.Fprintln(w, "return a single article in JSON")
	default:
		http.Error(w, "unsupported", http.StatusNotImplemented)
	}
}

// END OMIT

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "users")
}
