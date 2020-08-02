package main

import (
	"fmt"
	"sort"
)

func main() {
	var xi = []int{4, 6, 8, 3, 30, 21, 5, 6}
	var xs = []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	// both unsorted

	// sorting
	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
