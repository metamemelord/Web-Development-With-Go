package main

import (
	"html/template"
	"log"
	"net/http"
)

/*
	type Handler interface {
		ServerHTTP(ResponseWriter, *Request)
	}
*/

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Set-Cookie", "hehe=false")
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
