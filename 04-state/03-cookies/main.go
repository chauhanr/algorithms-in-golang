package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/expire", expire)
	http.HandleFunc("/read", read)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "value1",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN-CHECK BROWSER")
	fmt.Fprintln(w, "in chrome go to applicaiton dev tools")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Your Cookie: %s\n", c.Value)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	return
}
