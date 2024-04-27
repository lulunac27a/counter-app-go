package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	var count int = 0
	http.HandleFunc("/increase-counter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var counter int
			counter, err := strconv.Atoi(r.FormValue("counter"))
			count += counter
			log.Printf("Counter: %d", count)
			fmt.Fprintf(w, "%d", count)
			if err != nil {
				panic(err)
			}
		}
	})
	log.Print("Serving at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
