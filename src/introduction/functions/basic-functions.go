package main // Entry point

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
	fmt.Println()
	sayGreeting("Hi", "Bob")
	fmt.Println()
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name string) { // Can use comma delimited list for repeated types
	fmt.Println(greeting, name)
}

func sum(values ...int) int { // Variadic parameters + Return value (can also return value as a pointer)
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
