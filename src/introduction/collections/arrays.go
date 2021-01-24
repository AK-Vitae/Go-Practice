package main

import (
	"fmt"
)

func main() {
	var students [2]string //var name [size]type
	students[0] = "Jack"
	students[1] = "Jill"
	grades := [...]int{97, 85, 93} //[size]type{data} or [...]type{data}
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students)) //len()
	fmt.Printf("Grades: %v\n", grades)

	var identityMatrix = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("Identity Matrix: %v\n", identityMatrix)

	a := [...]int{1, 2, 3}
	b := a //A literal copy
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
