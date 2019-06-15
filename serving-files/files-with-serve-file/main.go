package main

import (
	"net/http"
)

type bytes []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(bytes("<img src=\"/doggo.jpg\" />"))
	})

	http.HandleFunc("/doggo.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "doggo.jpg")
	})

	http.ListenAndServe(":8080", nil)
}
