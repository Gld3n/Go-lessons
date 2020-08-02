package main

import "fmt"

func main() {
	far()

	// there's no identification
	func() {
		fmt.Println("Anon func executed.")
	}() // parenthesis to auto-execute the func

	func(x int) {
		fmt.Println("Number:", x)
	}(5) // in case of receiving params, they've to be passed into the parenthesis

}

func far() {
	fmt.Println("Far executed.")
}
