package main

import (
	"fmt"
)

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{1, 2, 3}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int))) // Need to type assert if you want to get the values

	you := person{}

	you.info = "Name"
	fmt.Println(you.info)
	you.returnInterfaceType()

	you.info = 40
	fmt.Println(you.info)
	you.returnInterfaceType()
}

type emptyInterface interface{}

type person struct {
	info interface{}
}

func (p person) returnInterfaceType() {
	switch p.info.(type) {
	case int:
		fmt.Println("info is an integer")
	case string:
		fmt.Println("info is a string")
	default:
		fmt.Println("idk bruv")
	}
}
