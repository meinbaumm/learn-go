// https://go.dev/doc/articles/wiki/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func maxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s Max!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", maxHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
