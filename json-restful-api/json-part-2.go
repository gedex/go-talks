package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name     string
	Username string
}

var user User

func main() {
	log.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}

// START OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&user); err != nil {
			msg := fmt.Sprintf("bad request: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
		}
		fmt.Fprintf(w, "Howdy, %s!\n", user.Username)
	case http.MethodGet:
		user = User{"Akeda Bagus", "gedex"}
		enc := json.NewEncoder(w)
		if err := enc.Encode(user); err != nil {
			msg := fmt.Sprintf("internal error: %v", err)
			http.Error(w, msg, http.StatusInternalServerError)
		}
	default:
		http.Error(w, "not supported", http.StatusNotImplemented)
	}
}

// END OMIT
