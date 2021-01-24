package main

import (
	"fmt"
)

// Structs can have different types of data
// Keyed by named fields
// Structs are a value type
type Car struct {
	quantity  int
	brandName string
	models    []string
}

func main() {
	aCar := Car{
		quantity:  50,
		brandName: "Toyota",
		models: []string{
			"Corolla",
			"Camry",
			"Sienna",
		},
	}
	fmt.Printf("Car 1: %v\n", aCar)
	fmt.Printf("Car 1: %v\n", aCar.models[1])
	// Dont need to include field names (Not Recommended as changes to struct syntax will cause errors)
	aCar2 := Car{
		50,
		"Honda",
		[]string{
			"Accord",
			"Civic",
		},
	}
	fmt.Printf("Car 2: %v\n", aCar2)
	// Anonymous Struct: Used in a short lived situation where you need to structure data
	aCar3 := struct {
		name string
	}{name: "Chevy"}
	aCar4 := &aCar3
	aCar4.name = "Tesla"
	fmt.Printf("Car 3: %v\n", aCar3)
	fmt.Printf("Car 4: %v\n", aCar4)
}
