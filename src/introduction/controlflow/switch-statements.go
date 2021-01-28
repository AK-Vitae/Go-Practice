package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Basic Switch syntax.
	// Breaks are implied
	switch 5 { // The case being tested is called a tag
	case 1, 5, 10: // Can test multiple cases but they have to be unique
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// Switch using an initializer and a semicolon separating from tag
	switch i := 2 + 3; i { // Can use an initializer
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")

	}

	// Switch with tag-less syntax: Case statements have full comparison operators
	rand.Seed(time.Now().UnixNano())
	j := rand.Intn(30)
	fmt.Printf("%v, ", j)
	switch {
	case j <= 10:
		fmt.Println("Less than or equal to ten")
		fallthrough
	case j <= 20:
		fmt.Println("Less than or equal to twenty")
	default:
		fmt.Println("Greater than twenty")
	}

	// Type Switch: Can be helpful when transforming data
	var i interface{} = [3]int{}
	switch i.(type) { // Tells Go to pull the underlying type used by the interface
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is a [3]int")
	default:
		fmt.Println("i is another type")
	}
}
