package main

import "fmt"

// pings is a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings is for receiving
// pongs is for sending
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
