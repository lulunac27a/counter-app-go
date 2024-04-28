package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fs := http.FileServer(http.Dir("./")) //use root directory ("/") for file server
	http.Handle("/", fs)                  //handle HTML requests to "/" (the root page)
	var count int = 0                     //initialize key count to 0
	var squareCount int = 0               //initalize square count to 0
	//Sends a POST request
	//w - HTTP response handler
	//r - HTTP request
	http.HandleFunc("/increase-counter", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost { //send POST request
			var counter int                                      //initialize counter variable to increase key count
			counter, err := strconv.Atoi(r.FormValue("counter")) //parse counter form variable to int
			count += counter                                     //increase count by counter
			squareCount += counter * counter                     //increase square count by counter times counter
			log.Printf("Counter: %d", count)                     //log the total key count
			log.Printf("Square Counter: %d", squareCount)        //log the square count
			fmt.Fprintf(w, `{"count": %d,
			"squareCount": %d}`, count, squareCount) //write POST response data
			if err != nil {
				panic(err) //stop if there is an error (error due to not a number when parsing to int)
			}
		}
	})
	log.Print("Serving at localhost:8080") //log when server starts
	http.ListenAndServe(":8080", nil)      //serve the web application at localhost:8080
}
