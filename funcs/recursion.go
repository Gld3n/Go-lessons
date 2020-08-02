package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println(f)

	f2 := cyFac(4)
	fmt.Println(f2)

}

// factorial with a recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

// factorial function with a -for- cycle
func cyFac(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total

}

// recursion is unnecesary and in some cases unreadable
// cycles as -for- can do the same job easier

// recursion defines -funcs- that calls itself repeatively.

// 4 * factorial(3)
// 3! = 3 * factorial(2)
// 2! = 2 * factorial(1)
// 1! = 1 * factorial(0)
// 0! = 1
