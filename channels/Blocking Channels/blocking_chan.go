package main

import (
	"fmt"
)

func main() {
	// unbuffered channel
	ca := make(chan int)

	// unbuffered channels can cause a block during run-time if
	// there's no goroutine waiting for its data

	ca <- 42
	fmt.Println(<-ca)
}
