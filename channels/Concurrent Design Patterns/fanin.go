package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		// get the first ten strings
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	// make a bidirectional channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			// put infinitely a formatted string in c with a random waiting time
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return c converting it to receive-only
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	// make a bidirectional channel
	c := make(chan string)
	go func() {
		for {
			// put infinitely the first input in c
			c <- <-input1
		}
	}()
	go func() {
		for {
			// put infinitely the second input in c
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
