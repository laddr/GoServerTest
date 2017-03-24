package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	c := make(chan int)
	go inc(c)
	fmt.Fprintf(w, strconv.Itoa(counter))
}

func inc(c chan int) {
	counter++
	c <- counter
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello There!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/increment", incrementCounter)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
