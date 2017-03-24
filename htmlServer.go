package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello There!")
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hi", hi)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
