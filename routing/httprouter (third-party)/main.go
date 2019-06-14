package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type bytes []byte

func main() {
	router := httprouter.New()
	router.GET("/hehe/*lol", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println(p[0].Key, p[0].Value)
		w.Write(bytes("hehe has dele"))
	})

	router.GET("/he/:ok", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println(p)
		w.Write(bytes("hehe has dele"))
	})

	http.ListenAndServe(":8080", router)
}
