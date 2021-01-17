package main

import (
	"fmt"
)

const (
	a = iota // Enumerated Constant
	b        // iota is inferred
	c
	d
	e
)

const (
	_  = iota // ignore first value by assigning to a blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	//Constants: Value must be calculable at compile time. PascalCase for exported constants and camelCase for internal constants
	const myConst int = 42 // Typed Constant
	fmt.Printf("Typed Constant: %v, %T\n", myConst, myConst)
	const myConst2 = 43 // Untyped Constant
	fmt.Printf("Untyped Constant: %v, %T\n", myConst2, myConst2)
	fmt.Printf("Enumerated Constant(iota): %v, %v, %v, %v, %v\n", a, b, c, d, e)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)

}
