package main

import (
	"encoding/json"
	"fmt"
)

// -Unmarshal- decodes JSON data to be usable by our program
func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoli", "Order": "Dasyuromorphia"}	
]`)
	type Animal struct {
		Name  string `json:"Name"`
		Order string `json:"Order"`
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i, v := range animals {
		fmt.Println("\nAnimal N°", i+1)
		fmt.Println("\t", v.Name, v.Order)
	}
}
