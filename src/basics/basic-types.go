package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false // Values are true or false, and zero value is false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//Variables declared without an explicit initial value are given their zero value
	fmt.Println("Zero Values")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//Bit operators
	c := 10 //1010
	v := 3  //0011
	fmt.Printf("Bit Operators: on %v and %v\n", c, v)
	fmt.Println("&", c&v)   // 0010
	fmt.Println("|", c|v)   // 1011
	fmt.Println("^", c^v)   // 1001
	fmt.Println("&^", c&^v) // 0100

	//Bit Shifting
	a := 8 //2^3
	fmt.Println("Bit Shifting")
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 1

	//Strings
	str := "this is a string"
	str2 := "this is also a string"
	x := []byte(str)                                           //Using byte slices
	fmt.Printf("Type: %T Value: %v\n", string(str[2]), str[2]) // need to use string() because strings are just bytes
	fmt.Printf("%v, %T\n", str+str2, str+str2)                 // String concatenation
	fmt.Printf("%v, %T\n", x, x)

	//Rune: Type alias for int32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
