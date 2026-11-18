package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()

		clean := cleanInput(line)
		fmt.Println("Your command was", clean[0])

	}
}

func cleanInput(text string) []string {
	dividedText := make([]string, 0)
	for _, word := range strings.Fields(text) {
		formattedWord := strings.ToLower(word)
		dividedText = append(dividedText, formattedWord)
	}

	return dividedText
}
