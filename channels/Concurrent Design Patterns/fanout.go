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

	go addint(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func addint(c chan int) {
	// put the index number in c 100 times
	for i := 0; i < 100; i++ {
		c <- i
	}
	// close the channel to avoid blocking the execution when
	// range is applied to c
	close(c)
}

// FAN OUT
func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	// cycle through c1 and launch a goroutine which puts in c2 what
	// expensiveJob returns
	for v := range c1 {
		wg.Add(1) // for each range and a goroutine
		go func(v2 int) {
			c2 <- expensiveJob(v2)
			wg.Done() // finish the goroutine
		}(v)
	}
	wg.Wait()
	// close c2 channel to prevent blocking execution when
	// applying range to it
	close(c2)
}

func expensiveJob(n int) int {
	// sets a random waiting time and returns a random integer
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
