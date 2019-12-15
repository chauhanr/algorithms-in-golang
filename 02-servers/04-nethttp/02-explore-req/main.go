package main

import (
	"html/template"
	"log"
	"net/http"
)

type server int

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
	var s server
	http.ListenAndServe(":9090", s)
}
