package main

import "fmt"

func main() {
	// buffered channel
	ca := make(chan int, 1)

	// since I got a buffered chan, I don't need other goroutine
	ca <- 42

	fmt.Println(<-ca)
}
