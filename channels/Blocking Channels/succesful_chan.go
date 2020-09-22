package main

import "fmt"

func main() {
	// unbuffered channel
	ca := make(chan int)

	// it's created another goroutine to solve the problem
	go func() {
		// data it's sent trough the channel with a goroutine
		ca <- 42
	}()

	// receive and show data from that chan
	fmt.Println(<-ca)
}
