package main

import "fmt"

func main() {
	summ := summ(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(summ)
}

// -...T- operator to say it's gonna take variadic params of some type
// -...T- must always be the final param
func summ(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	summ := 0
	for i, v := range x {
		summ += v
		fmt.Println("El valor en el indice", i, "suma", v, "al total, quedando", summ)

	}
	fmt.Println("El total es", summ)
	return summ

}
