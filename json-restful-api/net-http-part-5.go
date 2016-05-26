package main

import (
	"log"
	"net/http"
)

// START OMIT
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

// END OMIT

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

// START MAIN
func main() {
	handler := http.HandlerFunc(root)

	http.Handle("/", middlewareOne(middlewareTwo(handler)))
	http.ListenAndServe(":8080", nil)
}

// END MAIN
