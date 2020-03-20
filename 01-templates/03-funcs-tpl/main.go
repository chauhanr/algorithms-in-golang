package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	funcs := template.FuncMap{"title": strings.Title}

	tpl, err := template.New("base").Funcs(funcs).ParseFiles("funcs.tpl")
	if err != nil {
		log.Fatalln("Error parsing base tpl", err)
	}
	data := "the go programming language"

	err = tpl.ExecuteTemplate(os.Stdout, "funcs.tpl", data)
	if err != nil {
		log.Fatalln("error executing base template", err)
	}

}
