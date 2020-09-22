package main

import "fmt"

func main() {
	// a send-only channel
	ca := make(chan<- int, 2)

	ca <- 42
	ca <- 84

	// doing this is going to cause an error due to trying to receive
	// from an send-only channel
	fmt.Println(<-ca)
}
