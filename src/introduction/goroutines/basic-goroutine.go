package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)

	// Example 2
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye" // This portion runs before the goroutine
	time.Sleep(100 * time.Millisecond)

	// Example 3
	msg = "Hello"
	go func(msg string) { // Pass in argument
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
