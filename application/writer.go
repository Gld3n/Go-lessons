package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello world")
	fmt.Fprintln(os.Stdout, "Hello world")
	io.WriteString(os.Stdout, "Hello world")
}
