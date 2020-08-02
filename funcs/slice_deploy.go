package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...) // -...- to deploy the slice and give as argument the slice elements
	fmt.Println(x)
}

func sum(x ...int) int {
	fmt.Println(len(x), cap(x))
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("El valor en el indice", i, "suma", v, "al total, quedando", sum)

	}
	return sum

}
