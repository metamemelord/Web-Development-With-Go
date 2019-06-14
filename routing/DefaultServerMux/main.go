package main

import "net/http"

/*
	type handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/

type bytes []byte
type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes("{\"hehe\": \"lol\"}"))
}

func hehe(w http.ResponseWriter, r *http.Request) {
	w.Write(bytes("hehe has dele"))
}

func main() {
	hd := new(hotdog)
	http.Handle("/dog/", hd) // Registers the handler in DefaultServerMux
	http.Handle("/hehe", http.HandlerFunc(hehe))
	http.HandleFunc("/dog",
		func(w http.ResponseWriter, r *http.Request) { // Registers the function in DefaultServerMux
			w.Write(bytes("Hehe has dele"))
		})

	http.ListenAndServe(":8080", nil) // Handles using DefaultServeMux
}
