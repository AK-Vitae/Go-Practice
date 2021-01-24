package main

import "fmt"

func main() {
	// Data is passed by reference
	statePopulations := make(map[string]int) //One way to declare a map
	statePopulations = map[string]int{       // key-value
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])  //Using the key as variable or literal
	statePopulations["Georgia"] = 10310371 //Adding to map
	fmt.Println(statePopulations)          //Order of elements in the map aren't fixed
	delete(statePopulations, "Georgia")    //delete(map, key) = removes the key
	fmt.Println(statePopulations)
	_, ok := statePopulations["Oho"] //Used to check if key is in map
	fmt.Println(ok)
	fmt.Printf("Length of Map: %v\n", len(statePopulations))

}
