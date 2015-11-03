package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm served from %s!", r.Header.Get("X-Forwarded-For"))
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8484", nil)
}
