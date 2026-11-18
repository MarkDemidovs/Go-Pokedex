package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()

		clean := cleanInput(line)
		firstArgument := clean[0]

		command, ok := supportedCommands[firstArgument]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}

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
