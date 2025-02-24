package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}
