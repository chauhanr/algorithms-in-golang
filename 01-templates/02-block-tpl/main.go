package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	funcs := template.FuncMap{"join": strings.Join}

	tpl, err := template.New("base").Funcs(funcs).ParseFiles("block-master.tpl")
	if err != nil {
		log.Fatalln("Error parsing base tpl", err)
	}
	data := []string{"steve", "tony", "thor", "bruce", "clint", "natasha"}
	overlay, err := template.Must(tpl.Clone()).ParseFiles("overlay.tpl")
	if err != nil {
		log.Fatalln("error parsing overlay tpl", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "block-master.tpl", data)
	if err != nil {
		log.Fatalln("error executing base template", err)
	}

	err = overlay.ExecuteTemplate(os.Stdout, "block-master.tpl", data)
	if err != nil {
		log.Fatalln("error executing overlay template", err)
	}
}
