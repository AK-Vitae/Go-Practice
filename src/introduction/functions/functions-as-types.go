package main

import (
	"fmt"
)

func main() {
	func() { // Anonymous function
		fmt.Println("Hello Go!")
	}()

	for i := 0; i < 5; i++ {
		func(i int) { // Good practice is to pass in the variable from outer scope into the anonymous function
			fmt.Println(i)
		}(i)
	}

	f := func() { // Saving a function to variable
		fmt.Println("Hello Go!")
	}
	f()
}
