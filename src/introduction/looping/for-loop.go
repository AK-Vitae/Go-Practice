package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
	fmt.Println()

	for j, k := 0, 0; j < 5; j, k = j+1, k+2 {
		fmt.Println(j, k)
	}
	fmt.Println()

	// Break and Continue
	l := 0
	for {
		fmt.Println(l)
		l++
		if l == 5 {
			break
		}
	}
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			continue
		}
		fmt.Println(m)
	}
	fmt.Println()

	// Nested Loops and labels
Loop: // A label
	for n := 1; n <= 3; n++ {
		for o := 1; o <= 3; o++ {
			fmt.Println(n * o)
			if n*o >= 3 {
				break Loop
			}
		}
	}
}
