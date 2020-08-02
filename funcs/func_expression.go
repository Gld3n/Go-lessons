package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello world!")
	}

	f()

	g := func(y int) {
		fmt.Println("America was discovered in:", y)
	}

	g(1492)

}

// funcs are first-order citizens. That means they are also types
// as each other in the language(int, string, bool, interface, struct...)
// so they can be passed to a var
