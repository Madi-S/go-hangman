package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/fatih/color"
)

var (
	red     = color.New(color.FgRed).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	magenta = color.New(color.FgHiMagenta).SprintFunc()
)

var initialAttempts = 6

func main() {
	client := &http.Client{}
	word := getRandomWord(client)
	currentWordState := initializeCurrentWordState(word)

	attempts := initialAttempts
	guessedLetters := make(map[string]bool)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(blue("Welcome to Hangman!\n"))

	for attempts > 0 {
		fmt.Println(blue(fmt.Sprintf("Current word state: %s\n", strings.Join(currentWordState, " "))))
		fmt.Println(blue(fmt.Sprintf("Attempts left: %d\n", attempts)))

		userInput := getUserLoweredInput(scanner)

		if !isValidInput(userInput) {
			fmt.Println(red("Invalid input! Please enter a single letter."))
			continue
		}

		if guessedLetters[userInput] {
			fmt.Println(yellow("You've already guessed this letter."))
			continue
		}

		guessedLetters[userInput] = true

		if isGussedCorrectlyAndUpdateCurrentWordState(word, currentWordState, userInput) {
			fmt.Println(green("Good job, you've guessed the letter!"))
		} else {
			attempts -= 1
			fmt.Println(red("Incorrect guess."))
		}

		displayHangman(initialAttempts - attempts)

		if !strings.Contains(strings.Join(currentWordState, ""), "_") {
			fmt.Println(green(fmt.Sprintf("Congratulations you've gussed the word '%s'\n", word)))
			os.Exit(0)
		}

		fmt.Println("")
	}
	fmt.Println(red(fmt.Sprintf("You suck! The word was %s.\n", word)))
	os.Exit(1)
}

func initializeCurrentWordState(word string) []string {
	currentWordState := make([]string, len(word))
	for i := range currentWordState {
		currentWordState[i] = "_"
	}
	return currentWordState
}

func getUserLoweredInput(scanner *bufio.Scanner) string {
	fmt.Print(cyan("You're guess: "))
	scanner.Scan()
	return strings.ToLower(scanner.Text())
}

func isValidInput(input string) bool {
	return utf8.RuneCountInString(input) == 1 && unicode.IsLetter([]rune(input)[0])
}

func isGussedCorrectlyAndUpdateCurrentWordState(word string, currentWordState []string, letter string) bool {
	isGussed := false
	for i, char := range word {
		if string(char) == letter {
			currentWordState[i] = letter
			isGussed = true
		}
	}
	return isGussed
}

func displayHangman(incorrectGuesses int) {
	if incorrectGuesses >= 0 && incorrectGuesses < len(hangmanStates) {
		fmt.Println(magenta(hangmanStates[incorrectGuesses]))
	} else {
		fmt.Println(magenta(hangmanStates[len(hangmanStates)-1]))
	}
}
