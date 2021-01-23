package main

import (
	"fmt"
	"reflect" // Needed for tag
)

type Canine struct {
	Name   string `required max: "100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Canine{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
