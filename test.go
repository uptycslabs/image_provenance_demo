package main

import "fmt"

func sanitizeUrl(redir string) string {
	if len(redir) > 0 && redir[0] == '/' {
		return redir
	}
	return "/"
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
}
