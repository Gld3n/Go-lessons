package main

import (
	"context"
	"fmt"
)

func main() {
	// set a key and value to the context
	ctx := context.WithValue(context.Background(), "language", "Go")
	foo(ctx, "language") // search the key in ctx

	ctx = context.WithValue(ctx, "dog", "Gaston")
	foo(ctx, "dog")

	foo(ctx, "color")
}

func foo(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil { // whether there's a value
		fmt.Println("Value found:", v) // print the value found
		return
	}
	fmt.Println("Key not found(404):", s) // if there's no value
}
