package main

import "fmt"

func main() {
	// a receive-only channel
	ca := make(<-chan int, 2)

	// it's not valid sending data to a receive-only chan
	ca <- 3
	ca <- 6

	fmt.Println(<-ca)
}
