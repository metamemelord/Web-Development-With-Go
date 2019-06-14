package main

import "net/http"

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hehe\": \"lol\"}"))
}

func main() {
	hd := new(hotdog)
	mux := http.NewServeMux()
	mux.Handle("/dog", hd)  // Catch things ONLY when there's an exact match '/dog'
	mux.Handle("/dog/", hd) // Catch things down the line

	http.ListenAndServe(":8080", mux)
}
