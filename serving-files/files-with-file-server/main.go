package main

import (
	"io"
	"net/http"
)

type bytes []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "<img src=\"files/doggo.jpg\" />")
	})

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
