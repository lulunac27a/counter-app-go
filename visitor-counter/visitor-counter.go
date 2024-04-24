package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	count := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		log.Printf("Counter: %d", count)
		fmt.Fprintf(w, "Counter: %d", count)
	})
	http.ListenAndServe(":8080", nil)
}
