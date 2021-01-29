package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var ms2 *myStruct
	fmt.Println(ms2)
	ms = new(myStruct)
	fmt.Println(ms2)

	// How to access underlying field of a nil pointer
	var ms3 *myStruct
	ms3 = new(myStruct)
	(*ms3).foo = 42         // Set
	fmt.Println((*ms3).foo) // Get
	ms3.foo = 52            // Set without pointer or parentheses
	fmt.Println(ms3.foo)    // Get without pointer or parentheses

	// Slices vs Arrays
	a := []int{1, 2, 3} // Slice has a pointer to internal array that it is projecting
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)

	// Maps
	c := map[string]string{"foo": "bar", "baz": "buz"}
	d := c
	fmt.Println(c, d)
	c["foo"] = "qux"
	fmt.Println(c, d)
}

type myStruct struct {
	foo int
}
