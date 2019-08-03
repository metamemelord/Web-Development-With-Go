package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", redirectHandlerFunc)
	http.HandleFunc("/new", redirectHandlerFuncWithBuiltin)
	http.HandleFunc("/red", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Yeah, cool, redirected")
	})

	http.ListenAndServe(":8080", nil)
}

func redirectHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/red")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func redirectHandlerFuncWithBuiltin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/red", http.StatusTemporaryRedirect)
}
