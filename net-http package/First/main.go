package main

import (
	"fmt"
	"net/http"
)

/*
	type Handler interface {
		ServerHTTP(ResponseWriter, *Request)
	}
*/

type hehe float64

func (m *hehe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Some test lines")
}

func main() {
	var handler *hehe
	http.ListenAndServe(":8080", handler)
}
