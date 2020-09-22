package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Finished.")
}
