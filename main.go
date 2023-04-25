package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("activity: This is activity %s\n", r.URL.Path)
		fmt.Printf("log: This is normal log %s\n", r.URL.Path)
		fmt.Fprintf(w, "Hello - 2, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("activity: This is activity %s\n", r.URL.Path)
		fmt.Printf("log: This is normal log %s\n", r.URL.Path)
		fmt.Fprintf(w, "Hi - 2")
	})

	fmt.Println("Start.")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
