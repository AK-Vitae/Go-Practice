package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	for i := 0; i < len(s); i++ { // Basic Loop approach
		fmt.Println(s[i])
	}
	for k, v := range s { // For range loop
		fmt.Println(k, v)
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
	for _, v := range statePopulations {
		fmt.Println(v)
	}
}
