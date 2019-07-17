package main

import (
	"net/http"
	"os"
)

type bytes []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(bytes("<img src=\"/doggo.jpg\" />"))
	})

	http.HandleFunc("/doggo.jpg", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("doggo.jpg")
		if err != nil {
			http.Error(w, "Could not open the file", 404)
			return
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			http.Error(w, "Could not open the file", 404)
			return
		}
		http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
	})

	http.ListenAndServe(":8080", nil)
}
