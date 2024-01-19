package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

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

func doAuthReq(authReq *http.Request) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, _ := client.Do(authReq)
	return res
}

func main() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	doAuthReq(req)
	controller("nope")
	var p *int
	fmt.Printf("%v\n", *p)
}
