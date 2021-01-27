package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Defer: Used to delay execution of a statement until function exits
	// Useful to group "open" and "close" functions together (Need to be careful in loops)
	// Run in LIFO (last in, first out) order
	// Arguments evaluated at time defer is executed, not at time of called function execution
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// More complex example:
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Do not use this syntax when looping
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// Deferred statements take arguments at the time of the defer
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
