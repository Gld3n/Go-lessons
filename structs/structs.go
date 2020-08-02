package main

import "fmt"

//Def a -struct-
type person struct {
	name     string
	lastName string
	age      int
}

type secretAgent struct {
	person  // Embedded field
	license bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			name:     "James",
			lastName: "Bond",
			age:      32,
		},
		license: true,
	}

	p1 := person{
		name:     "Alejandro",
		lastName: "Jimenez",
		age:      15,
	}

	// -EmbeddedFields- are promoted so I can call -name-
	// atribute without saying "sa1.person.name"
	fmt.Println(sa1.name, sa1.lastName, sa1.age)
	fmt.Println(p1.name, p1.lastName, p1.age)

}
