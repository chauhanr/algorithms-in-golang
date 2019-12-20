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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="images/dew.jpg">`)
}
