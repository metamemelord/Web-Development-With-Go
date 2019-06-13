package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type person struct {
	Name string
	Age  int
}

func main() {
	tpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, person{"Gaurav", 22})
	if err != nil {
		log.Fatalln(err)
	}

	fm := template.FuncMap{
		"uc": strings.ToUpper,
		"dtform": func(t time.Time) string {
			return t.Format("05/09/1996")
		},
	}

	tpl = template.Must(template.New("template.range.html").Funcs(fm).ParseFiles("template.range.html"))
	// // Does not work
	// tpl, err = template.ParseFiles("template.range.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// tpl.Funcs(fm) Functions need to be attached before parsing the template
	err = tpl.Execute(os.Stdout, []string{"Hehe", "has", "Dele", "rinkiya", "ke", "pApA"})
	if err != nil {
		log.Fatalln(err)
	}

	// This section is not complete.
}
