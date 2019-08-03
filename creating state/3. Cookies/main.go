package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set", setCookieHandler)
	http.HandleFunc("/read", getCookieHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/set", http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(":8080", nil)
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "x-cookie",
		Value: "foobar",
	})
	w.Header().Set("content-type", "text/html")
	io.WriteString(w, "Written the cookie!")
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("x-cookie")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("content-type", "text/html")
	io.WriteString(w, fmt.Sprintf("Your cookie: %s=%s", c.Name, c.Value))
}
