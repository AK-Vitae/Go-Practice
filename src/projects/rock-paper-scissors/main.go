package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Welcome Message
	fmt.Println("Welcome to Rock, Paper, Scissors!")
	name, rounds := getPlayerInfo()

	for i := 1; i <= rounds; i++ {
		fmt.Printf("%s, now beginning round %d...\n", name, i)
	}
}

// Read in player's name and how many rounds of rock, paper, scissors they want to play
func getPlayerInfo() (string, int) {
	// Buffered Reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %s\n", name)

	fmt.Print("Enter the number of rounds you would like to play: ")
	var (
		rounds int
		err    error
	)
	for {
		roundsInput, _ := reader.ReadString('\n')
		roundsInput = strings.TrimSpace(roundsInput)
		rounds, err = strconv.Atoi(roundsInput)
		if err != nil {
			fmt.Printf("Enter a valid number: ")
		} else {
			fmt.Println("Got: " + strconv.Itoa(rounds))
			break
		}
	}
	return name, rounds
}
