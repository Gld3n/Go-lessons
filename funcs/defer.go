package main

import "fmt"

func main() {
	defer yup()
	yap()

}

func yup() {
	fmt.Println("yup")

}

func yap() {
	fmt.Println("yap")

}
