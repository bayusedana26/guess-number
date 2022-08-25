package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create secret number
	secret := getRandomNumber()

	// Get user input
	for matching := false; !matching; {
		guess := getUserInput()
		// fmt.Println(secret, guess)

		// Compare
		matching = isMatching(secret, guess)
		fmt.Println(matching)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Ooops your guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Ouch your guess is too small")
		return false
	} else {
		fmt.Println("You got it")
		return true
	}
}

func getUserInput() int {
	fmt.Print("Please input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}
	return input
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
