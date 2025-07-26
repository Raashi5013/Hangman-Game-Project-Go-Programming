Hangman-Game-Project-Go-Programming
This is a Hangman game built with Go, featuring themed word categories and robust input validation. Players choose from five themes each containing ten lowercase words. A random word from the selected theme is chosen, and the player must guess it letter by letter within a limited number of attempts.

The game provides a clean and interactive user experience. It validates each input to ensure it’s a single alphabetic character and prevents repeated guesses(valid and invalid). The current state of the word is displayed with correctly guessed letters revealed and remaining attempts shown after each guess.

The project uses core Go libraries like fmt, math/rand, time, regexp, and strings, demonstrating effective use of control flow, maps, and slices. After each round, the player is asked if they’d like to play again, allowing continuous gameplay without restarting the program.

