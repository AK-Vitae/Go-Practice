package main

import (
	"fmt"
	"strconv" // Used for type conversion of strings
)

//Visibility:
//lowercase first letter is for package scope
//uppercase first letter is to export
//variable can be scoped to a block by declaring it within a block and not at the package level

//Wrapping variables to remove redundant var
var (
	carMake  string  = "Audi"
	carModel string  = "R8"
	year     int     = 2013
	mileage  float32 = 9100.23
)

//Declaring at the package level: Can't use colon syntax
var l int = 500

func main() {
	fmt.Printf("Car: %v, %v, %v, %v\n", carMake, carModel, year, mileage)

	//Variable Declaration
	//var i int
	//i = 42
	var j int = 27
	//k := 99
	fmt.Printf("variable j: %v, %T\n", j, j)
	fmt.Printf("variable l: %v, %T\n", l, l)

	//Type Conversions: Need to explicitly convert. Not valid for strings.
	var a int = 12
	fmt.Printf("%v, %T\n", a, a)

	var b float32
	b = float32(a) //destinationType(variable)
	fmt.Printf("%v, %T\n", b, b)

	var c string
	c = strconv.Itoa(a)
	fmt.Printf("%v, %T\n", c, c)
}
