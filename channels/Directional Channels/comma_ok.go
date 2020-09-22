package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c) // close the channel when there's no values remaining
	}()

	// check if there's values in c and show the value
	v, ok := <-c
	fmt.Println(v, ok)

	// c has no values so it's closed. ok should print false
	v, ok = <-c
	fmt.Println(v, ok)
}
