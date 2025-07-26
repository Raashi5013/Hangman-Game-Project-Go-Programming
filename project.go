package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {
	playAgain := true

	for playAgain {
		playHangman()
		playAgain = askPlayAgain()
	}
}

func playHangman() {
	// Define themes and their corresponding word lists
	themes := map[string][]string{
		"animals":   {"lion", "elephant", "tiger", "giraffe", "zebra", "monkey", "penguin", "cheetah", "crocodile", "hippopotamus"},
		"fruits":    {"apple", "banana", "orange", "strawberry", "kiwi", "pineapple", "grape", "watermelon", "mango", "pear"},
		"countries": {"usa", "china", "india", "russia", "brazil", "japan", "germany", "france", "canada", "australia"},
		"colors":    {"red", "blue", "green", "yellow", "orange", "purple", "black", "white", "brown", "pink"},
		"sports":    {"football", "basketball", "tennis", "volleyball", "baseball", "golf", "soccer", "swimming", "cricket", "rugby"},
	}

	// Display themes to the player and let them choose
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Choose a theme:")
	for theme := range themes {
		fmt.Println(theme)
	}
	var chosenTheme string

	// Loop until a valid theme is entered
	for {
		fmt.Print("Theme: ")
		fmt.Scanln(&chosenTheme)

		// Validate chosen theme
		if _, themeExists := themes[chosenTheme]; themeExists {
			break
		} else {
			fmt.Println("Invalid theme. Please choose a valid theme.")
		}
	}

	// Select a random word from the chosen theme's list
	wordList := themes[chosenTheme]
	wordToGuess := wordList[rand.Intn(len(wordList))]

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Maximum number of attempts
	maxAttempts := 6

	// Current state of guessed letters
	guessedLetters := make([]bool, len(wordToGuess))

	// Guessed letters sets
	incorrectGuesses := make(map[string]bool)
	correctGuesses := make(map[string]bool)

	// Game loop
	for attemptsLeft := maxAttempts; attemptsLeft > 0; {
		// Display current state of word with blanks for unguessed letters
		displayWord(wordToGuess, guessedLetters)
        
        // Show attempts left
		fmt.Printf("Attempts left: %d\n", attemptsLeft)

		// Get user input
		guess := getUserInput()

		// Check if the guessed letter is already guessed
		if incorrectGuesses[guess] || correctGuesses[guess] {
			if correctGuesses[guess] {
				fmt.Printf("Letter %s has already been guessed correctly.\n", guess)
			} else {
				fmt.Printf("Letter %s has already been guessed incorrectly.\n", guess)
			}
			continue
		}

		// Check if the guessed letter is in the word
		if strings.Contains(wordToGuess, guess) {
			// Update guessedLetters
			for i, letter := range wordToGuess {
				if string(letter) == guess {
					guessedLetters[i] = true
				}
			}

			// Check if all letters have been guessed
			if allLettersGuessed(guessedLetters) {
				displayWord(wordToGuess, guessedLetters)
				fmt.Println("Congratulations! You guessed the word!")
				return
			}
			fmt.Printf("Correct guess! Letter %s is in the word.\n", guess)
			correctGuesses[guess] = true
		} else {
			// Guessed letter is not in the word
			fmt.Println("Incorrect guess!")
			attemptsLeft--
			incorrectGuesses[guess] = true
		}
	}

	fmt.Println("Out of attempts! The word was:", wordToGuess)
}

func displayWord(word string, guessedLetters []bool) {
	for i, letter := range word {
		if guessedLetters[i] {
			fmt.Printf("%c ", letter)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func allLettersGuessed(guessedLetters []bool) bool {
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}

func askPlayAgain() bool {
	var playAgain string
	fmt.Print("Do you want to play again? (yes/no): ")
	fmt.Scanln(&playAgain)
	return strings.ToLower(playAgain) == "yes"
}

func getUserInput() string {
	var guess string
	for {
		fmt.Print("Guess a letter: ")
		fmt.Scanln(&guess)

		// Validate user input
		if isValidInput(guess) {
			break
		} else {
			fmt.Println("Invalid input. Please enter a single alphabetic character.")
		}
	}
	return strings.ToLower(guess)
}

func isValidInput(input string) bool {
	// Check if input is a single character
	if len(input) != 1 {
		return false
	}
	// Check if input is alphabetic
	matched, _ := regexp.MatchString("^[a-zA-Z]*$", input)
	return matched
}
