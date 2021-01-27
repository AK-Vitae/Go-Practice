package main

import (
	"fmt"
	"log"
)

func main() {
	// Panic: Occur when program cannot continue at all
	// Use for unrecoverable events but don't use when a file can't be opened, unless it is critical
	// Order of execution: functions -> deferred statements -> panic statements -> return value
	// Recover: Used to recover from panic and only useful in deferred functions
	// Current functions will not attempt to continue, but higher function in call stack will
	fmt.Println("start")
	defer func() { // Anonymous function
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something bad happened") // Error handling; Panics happen after deferred statements
	fmt.Println("end")              // This statement won't run
}
