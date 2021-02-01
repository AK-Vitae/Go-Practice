package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet() // Method invocation
	g.greetPointer()
	fmt.Println("New name is called", g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // Copy of object is passed in "Object Receiver"
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) greetPointer() { // Object is directly passed in "Pointer Receiver"
	fmt.Println(g.greeting, g.name)
	g.name = "Java"
}
