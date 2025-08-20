package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", LocalHost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func LocalHost(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello, ", name)
}
