package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", somefunc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func somefunc(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session-id")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		c = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
			MaxAge:   10,
		}
		http.SetCookie(w, c)
	}
	fmt.Println(c)
}
