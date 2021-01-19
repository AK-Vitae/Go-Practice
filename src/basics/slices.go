package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3} //[]type{data}
	b := a              //A reference to original data
	b[1] = 5
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("Length: %v\n", len(a))     //len()
	fmt.Printf("Capacity: %v\n\n", cap(a)) //cap()

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]   //slice of all elements
	e := c[3:]  //slice from 4th elements to end
	f := c[:6]  //slice first 6 elements
	g := c[3:6] //slice the 4th, 5th, and 6th elements.
	//First number is inclusive while second number is exclusive
	//From index 3 up to but not including index 6
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("g: %v\n\n", g)

	h := make([]int, 3, 100) //make(type, length, capacity)
	fmt.Printf("h: %v\n", h)
	fmt.Printf("Length: %v\n", len(h))
	fmt.Printf("Capacity: %v\n\n", cap(h))

	var i []int
	fmt.Printf("i: %v\n", i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	i = append(i, 1)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	i = append(i, []int{2, 3, 4, 5}...) //concatenate slices
	fmt.Printf("i: %v\n", i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n\n", cap(i))

	//j := i[1:] //remove first element
	//k := i[:len(i)-1] //remove last element
	l := append(i[:2], i[3:]...) //remove element from the middle
	fmt.Printf("i: %v\n", i)
	//fmt.Printf("j: %v\n", j)
	//fmt.Printf("k: %v\n", k)
	fmt.Printf("l: %v\n", l)
}
