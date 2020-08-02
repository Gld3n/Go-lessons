package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	license bool
}

// func (r receptor) identificator(params) {code}
func (sa secretAgent) greet() {
	fmt.Println("Hola, soy el agente", sa.first, sa.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Robert",
			last:  "Vale",
		},
		license: true,
	}
	fmt.Println(sa1)
	sa1.greet()
}
