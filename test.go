package main

import (
	"fmt"
	"log"
	"net/http"
)

func sanitizeUrl(redir string) string {
	if len(redir) > 0 && redir[0] == '/' {
		return redir
	}
	return "/"
}

func serve() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.Form.Get("user")
		pw := r.Form.Get("password")

		log.Printf("Registering new user %s with password %s.\n", user, pw)
	})
	http.ListenAndServe(":80", nil)
}

func controller(msg string) {
	switch {
	case msg == "start":
		fmt.Printf("start")
	case msg == "start":
		fmt.Printf("stop")

	default:
		panic("Message not understood.")
	}
}

func main() {
	u := "/abc/def"
	sanitizeUrl(u)
	controller("nope")
	var p *int
	fmt.Printf("%v\n", *p)
	serve()
}
