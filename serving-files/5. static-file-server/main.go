package main

import (
	"net/http"
)

type bytes []byte

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
