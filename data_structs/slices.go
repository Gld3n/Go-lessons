package main

import "fmt"

func main() {
	// array (don't use. Just a constructor).
	// y := [5]int

	// composite literal
	x := []int{3, 6, 9}

	//------------------------------------------------

	// -for range- clause to iterate over an slice (data structure)
	for i, v := range x {
		fmt.Println(i, v)

	}

	//------------------------------------------------

	// -append- function
	fmt.Println(x[0:2])
	x = append(x, 44, 23, 34, 941)
	fmt.Println(x)

	// divide slices and eliminate elements
	x = append(x[:2], x[5:]...)
	fmt.Println(x)

	//------------------------------------------------

	// -make- function to build data structures better
	slice := make([]int, 10, 10)
	fmt.Println(slice, len(slice), cap(slice))

	// add elements one by one
	slice[6] = 3
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 22)
	slice = append(slice, 26)
	fmt.Println(slice, len(slice), cap(slice))
	// as result: cap == 20 and len == 12

	// an slice is dynamic, so it can change its lenght
	// but you can't set a value off the range:
	// 	<< slice[11] = "something" >>
	// it would cause an error, and it's better to use -append-
	// but when the subyacent array of the slice reach its max capacity
	// it throws away the slice and create another from scrath
	// duplicating the capacity of the array
	// that makes the script to be slower and it's a bad practice
	// it's better to know or estimate the max capacity
	// in order to optimize the script and eviting
	// the subprocess overwork

	//------------------------------------------------

	// multidimensional slices
	rv := []string{"Robert", "Vale.", "Basket"}
	vs := []string{"Valeria", "Serrano.", "Videogames"}

	multi := [][]string{rv, vs}
	fmt.Println("\n", multi)

	// -----------------------------------------------

	// map (equivalent to a dict of Python)
	m := map[string]int{
		"Iron Man":    45,
		"Doc Strange": 56,
		"Spider Man":  16,
	}
	fmt.Println("\n", m)

	// access to a specific value
	fmt.Println(m["Iron Man"])

	// language comma ok
	v, ok := m["Robert"]
	fmt.Println(v, ok)

	// ok evaluates to a bool, trying to find the key in the map
	// << if ok == true >>
	if v, ok := m["Robert"]; ok {
		fmt.Println("Won't print.")
		fmt.Println(v, ok)

	} else if v, ok := m["Spider Man"]; ok {
		fmt.Println("Spidey")
		fmt.Println(v, ok)

	}

	// assign values to a map
	m["Robert"] = 16
	fmt.Println(m)

	// iterating trough a map with a -for clause-
	for k, v := range m {
		fmt.Println(k, v)

	}

	// delete a map element
	// delete(m, "Spider Man")
	// fmt.Println(m)

	m["Spider Man"] = 18

	if v, ok := m["Spider Man"]; ok {
		fmt.Println("ok")
		fmt.Println(v)
		delete(m, "Spider Man")

	}
	fmt.Println(m)

}
