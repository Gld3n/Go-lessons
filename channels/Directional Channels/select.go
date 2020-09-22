package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	exit := make(chan int)

	// send
	go func(e, o, ex chan<- int) {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				e <- i
			} else {
				o <- i
			}
		}
		ex <- 0

	}(even, odd, exit)

	// receive
	func(e, o, ex <-chan int) {
		for {
			select {
			// in case there's a value:
			case v := <-e:
				fmt.Println("Even:", v)
			case v := <-o:
				fmt.Println("Odd:", v)
			case v := <-ex:
				fmt.Println("Exit:", v)
				return // break the infinite loop
			}
		}
	}(even, odd, exit)

	fmt.Println("Finished.")
}
