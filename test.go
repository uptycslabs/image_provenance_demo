package main

import "fmt"

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
	controller("nope")
	var p *int
	fmt.Printf("%v\n", *p)
}

