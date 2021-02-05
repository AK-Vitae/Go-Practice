package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // Strongly typed channel that takes only int type
	wg.Add(2)
	go func(ch <-chan int) { // Receiving goroutine/Receive only channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // Sending goroutine/Send only channel
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
