package main

import "fmt"

// interfaces & polimorfism
// interfaces let us define behaviors
type human interface {
	greet() // -greet- is a secretAgentt's method, so sA is also the interface type
}

type personn struct {
	first string
	last  string
}

type secretAgentt struct {
	personn
	license bool
}

func (s secretAgentt) greet() {
	fmt.Println("Hola, soy", s.first, s.last, ". Un agente secreto.")

}

func (p personn) greet() {
	fmt.Println("Hola, soy", p.first, p.last)
}

func bar(h human) {
	switch h.(type) {
	case personn:
		fmt.Println("I've been passed to bar func (personn).", h.(personn).first)

	case secretAgentt:
		fmt.Println("I've been passed to bar func (secretAgentt).", h.(secretAgentt).first)

	}
	fmt.Println("I've been passed to bar func.", h)
}

func main() {
	sa1 := secretAgentt{
		personn: personn{
			first: "Robert",
			last:  "Vale",
		},
		license: true,
	}

	sa2 := secretAgentt{
		personn: personn{
			first: "Alejandro",
			last:  "Jimenez",
		},
		license: false,
	}

	p := personn{
		first: "Valeria",
		last:  "Serrano",
	}

	fmt.Println(sa1)
	sa1.greet()
	p.greet()

	bar(sa1)
	bar(sa2)
	bar(p)

}
