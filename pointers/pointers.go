package main

import "fmt"

func main() {
	a := 3
	fmt.Println(a)  // 3
	fmt.Println(&a) // 0x1189c080

	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int

	var b *int = &a // 0x1189c080
	fmt.Println(b)  // 0x1189c080
	fmt.Println(*b) // 3

	*b = 2020
	fmt.Println(*b)

}

// pointers are references to memory slots
