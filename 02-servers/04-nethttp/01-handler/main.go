package main

import (
	"fmt"
	"net/http"
)

type server int

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Writing to the response body.")
}

func main() {
	var s server
	http.ListenAndServe(":9090", s)

}
