package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // finish when main is over
	for n := range gen(ctx) {
		fmt.Println(n)
		// get only the first 5 values and exit main
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int) // create a destiny channel
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // retornando para que no se fuge la gorutina
			// pass the value in n to destiny and add 1 to it
			case dst <- n:
				n++
			}
		}
	}()
	return dst // return destiny as a receive-only channel
}
