package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var s string = `somepass`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPass := `somepass`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		fmt.Println("You couldn't log in")
		return
	}
	fmt.Println("Logged.")

}
