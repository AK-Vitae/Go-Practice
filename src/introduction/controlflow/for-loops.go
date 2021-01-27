package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ { // Basic for loop
		fmt.Println(i)
	}
	fmt.Println()

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println()

	// Using break and continue
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Nested loop and labels
Loop: // A label
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

}
