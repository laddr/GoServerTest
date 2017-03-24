package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func maxClients(h http.Handler, n int) http.Handler{
	sema := make(chan struct{}, n)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sema <- struct{}{}
		defer func() { <- sema}()

			h.ServeHTTP(w,r)
	})
}

func main() {
	handler := http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.Handle("/",maxClients(handler,1))
	

	log.Fatal(http.ListenAndServe(":8081", nil))
}
