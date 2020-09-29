package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// returns withDeadline(parent, time.Now().Add(timeout))
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
