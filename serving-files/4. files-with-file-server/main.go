package main

import (
	"io"
	"net/http"
)

type bytes []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, "<img src=\"assets/doggo.jpg\" />")
	})

	http.Handle("/assets/", http.FileServer(http.Dir(".")))

	// General result of http.Handle(/uri, http.FileServer(http.Dir(fs_path))) is:
	// Serves fs_path + uri/* at uri/*

	// http.Handle("/assets/", http.FileServer(http.Dir("."))) - Serves dir ./assets/* at ./assets/*
	// http.Handle("/assets/", http.FileServer(http.Dir("./assets"))) - Serves dir ./assets/assets/* at ./assets/*
	// http.Handle("/", http.FileServer(http.Dir("./assets"))) - Serves dir ./assets at /
	// http.Handle("/", http.FileServer(http.Dir("."))) - Serves current directory at /

	// http.StripPrefix() strips the URI
	// http.Handle("/test/", http.StripPrefix("/test", http.FileServer(http.Dir("./assets")))) Serves ./assets/* at ./test/*

	http.ListenAndServe(":8080", nil) // http.FileServer is a handler. It can be sent into ListenAndServe directly.
}
