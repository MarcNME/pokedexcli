package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		command, ok := getCommands()[cleanedInput[0]]

		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Printf("unexpected Error: %v", err)
			}
		}

	}
}

func cleanInput(text string) []string {
	words := make([]string, 0)
	var tmp string
	for _, r := range []rune(text) {
		if unicode.IsSpace(r) {
			if len(tmp) > 0 {
				words = append(words, tmp)
				tmp = ""
			}
		} else {
			tmp += string(unicode.ToLower(r))
		}
	}

	if len(tmp) > 0 {
		words = append(words, tmp)
	}

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
	}
}
