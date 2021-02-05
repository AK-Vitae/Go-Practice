package main

import (
	"fmt"
	"sync"
)

var group = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // A buffer is added allowing for intake of multiple integers
	group.Add(2)
	go func(ch <-chan int) { // Receiving goroutine/Receive only channel
		for i := range ch {
			fmt.Println(i)
		}
		group.Done()
	}(ch)
	go func(ch chan<- int) { // Sending goroutine/Send only channel
		ch <- 42
		ch <- 27
		close(ch) // Close the channel
		group.Done()
	}(ch)
	group.Wait()
}
