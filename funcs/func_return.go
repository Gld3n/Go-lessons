package main

import "fmt"

func main() {
	// x := foar()
	// fmt.Printf("%T\n", x)
	fmt.Println(bar()())

}

func foar() func() int {
	return func() int {
		return 1492
	}
}
