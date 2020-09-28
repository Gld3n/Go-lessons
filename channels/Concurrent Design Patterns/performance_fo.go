package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go addInt(c1)

	go fanOutIn(c1, c2)

	// extract the values inside c2
	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Finishing.")
}

func addInt(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	// set a limit of goroutines
	const goroutines = 10
	wg.Add(goroutines)

	// launch a goroutine while i < goroutines(10)
	for i := 0; i < goroutines; i++ {
		go func() {
			// for each goroutine we iterate through c1
			for v := range c1 {
				// a closure which does the expensiveJob and
				// pass the returned value to c2 channel
				func(v2 int) {
					c2 <- expensiveJob(v2)
				}(v) //receiving v as argument(each element inside c1)
			}
			wg.Done()
		}()
	}
	wg.Wait() // wait for the goroutines to finish
	close(c2) // close the channel to prevent execution blocking due to range
}

func expensiveJob(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
