package main

import "fmt"

func main() {
	foo()
	bark("Robert")      // passing arguments
	s1 := woo("Robert") // instance of a func
	fmt.Println(s1)
	x, y := saludar("Robert", "Vale")
	fmt.Println(x, y)
}

// syntax: func (r receptor) identificator(parameters) (return(s)) { code }
func foo() {
	fmt.Println("Hello from foo.")

}

func bark(s string) {
	fmt.Println("Hello,", s)

}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo, ", st)
	// as -woo- returns a string, the -return- keyword is used with that purpose

}

func saludar(n, ln string) (string, bool) {
	p := fmt.Sprint(n, " ", ln, " dice 'Hola'")
	q := true
	return p, q

}
