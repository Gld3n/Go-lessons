package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	// set a time limit to the context
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel() // cancel when main is over

	select {
	// in case there's something in time.After()
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	// in case there's something in ctx.Done()
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // print the context error
	}

}
