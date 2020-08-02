package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Robert",
		last:  "Vale",
		age:   15,
	}
	fmt.Println(p1)

}
