package main

import (
	"fmt"
	"math"
)

func main() {
	// Basic Syntax
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// Initializer with a boolean block
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// Comparison Operators <, >, <=,>=, !=
	number := 50
	guess := 50
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		} else if guess > number {
			fmt.Println("Too High")
		} else if guess == number {
			fmt.Println("You got it!")
		}
	}

	// Using error parameters
	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
