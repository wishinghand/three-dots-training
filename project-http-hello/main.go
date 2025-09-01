package main

import (
	"fmt"
	"log"
	"net/http"
)

var calls []string
var stats map[string]int

func main() {
	stats = map[string]int{}
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
	
	calls = append(calls, name)
	stats[name]++

	fmt.Fprint(w, "Hello, ", name)
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
}
