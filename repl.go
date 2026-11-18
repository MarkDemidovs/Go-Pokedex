package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	commands map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()

		clean := cleanInput(line)

		if len(clean) == 0 {
			continue
		}

		firstArgument := clean[0]

		command, ok := cfg.commands[firstArgument]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
