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
	testTimestamp(1)
	controller("nope")
	var p *int
	fmt.Printf("%v\n", *p)
}
