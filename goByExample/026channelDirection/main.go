package main

import "fmt"

// ping only accepts a channel for sending values. It would be a compile-time
// error to try to receive on this channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong accepts one channel for receives(pings) and a second for sends(pongs)
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
