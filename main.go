package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!") 
	time.Sleep(3 * time.Second)
	fmt.Println("I'm thinking of a number between 1 and 100...")
	time.Sleep(3 * time.Second)
	
	fmt.Println("You have 5 chances to guess the correct number.")
	time.Sleep(3 * time.Second)
	
	fmt.Println("")
	
	fmt.Println("Please select the difficulty level:")
	time.Sleep(2 * time.Second)
	
	fmt.Println("")
	
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println("")
	fmt.Print("Enter your choice: ")

	var choice int
	var chances int

	fmt.Scanln(&choice)
	switch choice {
	case 1:
		chances = 10 
		fmt.Println("Great! You have selected the Easy difficulty level.")
	case 2:
		chances = 5 
		fmt.Println("Great! You have selected the Medium difficulty level.")
	case 3:
		chances = 3 
		fmt.Println("Great! You have selected the Hard difficulty level.")
	default:
		chances = 5 //default: medium
		fmt.Println("Invalid choice. Defaulting to Medium difficulty.")
	}

	fmt.Println("Let's start the game!")

	secNum := rand.IntN(100) + 1

	for attempts := 1; attempts <= chances; attempts++ {
		var guess int
		fmt.Print("Enter your Guess: ")
		_, err := fmt.Scanln(&guess)
		
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			var discard string
			fmt.Scanln(&discard)
			attempts--
			continue
		}

		

		if guess == secNum {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		}
		
		if attempts < chances {
			if guess < secNum {
				fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
			} else {
				fmt.Printf("Incorrect! The number is less than %d.\n", guess)
			}
		}
	}
	fmt.Printf("Game Over! You ran out of chances. The correct number was %d.\n", secNum)
}
