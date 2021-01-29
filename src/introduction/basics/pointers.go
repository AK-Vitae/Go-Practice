package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a // * before a type = pointer of that type; & = address of operator
	a = 27
	fmt.Println(a, b)  // b is holding the memory location of a
	fmt.Println(a, *b) // * = dereferencing operator
	*b = 14
	fmt.Println(a, *b)

	// Pointer Arithmetic: Need to use unsafe package
	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)
}
