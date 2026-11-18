package main

import "strings"

func cleanInput(text string) []string {
	dividedText := make([]string, 0)
	for _, word := range strings.Fields(text) {
		formattedWord := strings.ToLower(word)
		dividedText = append(dividedText, formattedWord)
	}

	return dividedText
}
