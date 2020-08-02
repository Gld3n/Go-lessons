package main

import "fmt"

func main() {
	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a := adition(is...)
	fmt.Println(a)

	s1 := pairsAdition(adition, is...)
	fmt.Println("El total de pares es:", s1)

	s2 := unpairsAdition(adition, is...)
	fmt.Println("El total de impares es:", s2)
}

func adition(xi ...int) int {
	// fmt.Printf("%T\n", xi)
	total := 0

	for _, v := range xi {
		total += v
	}
	return total

}

//		 (func adition(is...) int
func pairsAdition(f func(yi ...int) int, vi ...int) int {
	var y []int // the slice I will append to in case of pair

	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v) // appends if pair
		}
	}
	return f(y...) // returns sum(is...) if passed

}

func unpairsAdition(f func(yi ...int) int, vi ...int) int {
	var y []int

	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)

}
