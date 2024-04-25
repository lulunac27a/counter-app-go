package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var count = 0 //initialize visitor count to 0
	// Sends a request to "/" (the root page)
	// w - HTTP response handler
	// r - HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++                              //increase count when request is made
		log.Printf("Counter: %d", count)     //log the visitor count
		fmt.Fprintf(w, "Counter: %d", count) //print the visitor count
	})
	log.Print("Serving at localhost:8080") //log when server starts
	http.ListenAndServe(":8080", nil)      //serve the web application at localhost:8080
}
