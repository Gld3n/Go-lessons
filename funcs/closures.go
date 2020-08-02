package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	for i := 0; i < 6; i++ {
		fmt.Println(a())
	}

	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// code in curly braces -{}- is limited to that scope, such as creation of variables.
// limitations of scope is called -closure-.
