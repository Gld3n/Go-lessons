package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// -Marshal- converts our objects in JSON format
func main() {
	type ColorGroup struct {
		// capitalization is important when we want to convert
		// our -struct- into -json-
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Red",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
