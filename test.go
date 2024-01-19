package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	_ "github.com/consensys/gnark"
)

func sanitizeUrl(redir string) string {
	if len(redir) > 0 && redir[0] == '/' {
		return redir
	}
	u, err := url.Parse(redir)
	if err != nil || u.Scheme == "javascript" {
		return "about:blank"
	}
	return redir

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

type Timestamp int

func (t Timestamp) addDays(d int) Timestamp {
	return Timestamp(int(t) + d*24*3600)
}

func testTimestamp(t Timestamp) {
	fmt.Printf("Before: %s\n", t)
	t.addDays(7)
	fmt.Printf("After: %s\n", t)
}

func main() {
	u := "/abc/def"
	sanitizeUrl(u)
	testTimestamp(1)
	controller("nope")
	var p *int
	fmt.Printf("%v\n", *p)
	serve()
}
