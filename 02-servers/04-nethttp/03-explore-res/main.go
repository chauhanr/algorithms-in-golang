package main

import (
	"fmt"
	"net/http"
)

type server int

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Key", "This is from key")
	w.Header().Set("Content-Key", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var s server
	http.ListenAndServe(":9090", s)
}
