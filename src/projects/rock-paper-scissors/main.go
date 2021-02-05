package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to Rock, Paper, Scissors!")
	var (
		name   string
		rounds int
	)
	name, rounds = getPlayerInfo()
	playGame(name, rounds)
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
			break
		}
	}
	return name, rounds
}

// Game logic: paper beats rock, rock beats scissors, and scissors beats paper
func playGame(name string, rounds int) {
	reader := bufio.NewReader(os.Stdin)
	playerScore, computerScore := 0, 0

	for i := 1; i <= rounds; i++ {
		fmt.Printf("%s, now beginning round %d...\n", name, i)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Scores:\n\t %s: %d, Computer: %d\n\n", name, playerScore, computerScore)
		isDraw := false // If round results in a draw, repeat the round until a winner is decided

		for {
			if isDraw == true {
				fmt.Printf("Continuing round %d...\n", i)
				isDraw = false
			} else {
				fmt.Print("What is your move? To make a move, enter rock, paper, or scissors: ")
				playerMove, _ := reader.ReadString('\n')
				playerMove = strings.TrimSpace(playerMove)

				if playerMove != "rock" && playerMove != "paper" && playerMove != "scissors" { // Check for proper input
					fmt.Println("Please enter a valid move")
				} else {
					moveChoices := []string{"rock", "paper", "scissors"}
					index := rand.Intn(3) // Make computer's choice pseudorandom
					computerMove := moveChoices[index]
					fmt.Println("Computer picked:", computerMove)
					time.Sleep(time.Millisecond * 750)

					if playerMove == computerMove {
						fmt.Print("It's a tie!\n\n")
						isDraw = true
						continue
					} else if (playerMove == "rock" && computerMove == "scissors") || (playerMove == "scissors" && computerMove == "paper") || (playerMove == "paper" && computerMove == "rock") {
						fmt.Print("You won the round!\n\n")
						playerScore++
					} else {
						fmt.Print("You lost the round.\n\n")
						computerScore++
					}
					break
				}
			}
		}
	}

	fmt.Printf("Game Over:\n\t Your Score: %d\n\t Computer Score: %d\n", playerScore, computerScore)
	if playerScore > computerScore {
		fmt.Println("YOU WIN!")
	} else if playerScore < computerScore {
		fmt.Println("YOU LOSE.")
	} else {
		fmt.Println("TIE GAME!")
	}
}
