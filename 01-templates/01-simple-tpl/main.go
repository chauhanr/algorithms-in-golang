package main

import (
	"log"
	"os"
	"text/template"
)

type Inventory struct {
	Count    int
	Material string
}

func main() {
	in := Inventory{17, "Wool"}
	tpl, err := template.ParseFiles("simple.tpl", "withex.tpl", "nestedt.tpl")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "simple.tpl", in)
	if err != nil {
		log.Fatalln("Error executing template", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "withex.tpl", nil)
	if err != nil {
		log.Fatalln("Error executing template", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "nestedt.tpl", nil)
	if err != nil {
		log.Fatalln("Error executing template", err)
	}

}
