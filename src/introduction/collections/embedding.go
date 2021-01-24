package main

import "fmt"

// Go does not have inheritance but can use composition via embedding
// Struct(s) "has/have" another struct
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // embed the Animal struct in the Bird struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// When declaring using literal syntax need to be aware that embedding is being used
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	// When you are manipulating from the outside you don't need to know that embedding is being used
	c := Bird{}
	c.Name = "Emu"
	c.Origin = "Australia"
	c.SpeedKPH = 48
	c.CanFly = false

	fmt.Println(b)
	fmt.Println(c)
}
