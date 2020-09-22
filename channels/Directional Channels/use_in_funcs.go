package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func(ch chan<- int) {
		c <- 42
	}(c)

	// receive
	func(ch <-chan int) {
		fmt.Println(<-ch)
	}(c)

	fmt.Println("Finished.")
}
