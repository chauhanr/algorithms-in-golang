package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/dew", dew)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}

func dew(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, "Something to do: "+v)
}
