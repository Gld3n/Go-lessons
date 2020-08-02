package main

import (
	"fmt"
	"sort"
)

type Human struct {
	Name string
	Age  int
}

type ByName []Human

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	p1 := Human{"Robert", 15}
	p2 := Human{"Valeria", 15}
	p3 := Human{"Jose", 11}
	p4 := Human{"Alejandro", 15}

	people := []Human{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
