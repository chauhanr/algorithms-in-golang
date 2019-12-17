package main

import (
	"io"
	"net/http"
)

type serveDog int

func (d serveDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog, dog, dog")
}

type serveCat int

func (d serveCat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kitty, kitty, kitty")
}

func main() {
	var d serveDog
	var c serveCat

	mux := http.NewServeMux()
	// /dog/ will map all other pathat like /dog/some /dog/some/other etc.
	mux.Handle("/dog/", d)
	// /cat will only match /cat and all other paths like /cat/some, /cat/some/other will be
	// 404.
	mux.Handle("/cat", c)

	http.ListenAndServe(":9090", mux)

}
