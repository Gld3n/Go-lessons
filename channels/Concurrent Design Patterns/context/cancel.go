package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// get a Context type and a CancelFunc
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error checking 1:", ctx.Err())
	fmt.Println("Num Goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for { // infinitely do a select-case
			select {
			// in case there's a value in the chan <-ctx.Done()
			// which is passed by cancel() when it's called
			case <-ctx.Done():
				return
			default:
				// by default n is increased by 1
				time.Sleep(time.Millisecond * 200)
				// for each 200 miliseconds it's going to print n(the process working)
				fmt.Println("", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error checking 2:", ctx.Err())
	fmt.Println("Num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("Finishing context.")
	cancel() // cancel the process and pass a value to ctx.Done()
	fmt.Println("Context canceled.")

	time.Sleep(time.Second * 2)
	fmt.Println("Error checking 3:", ctx.Err())
	fmt.Println("Num goroutines 3:", runtime.NumGoroutine())
}
