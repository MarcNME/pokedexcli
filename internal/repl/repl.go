package repl

import (
	"bufio"
	"fmt"
	"github.com/marc-enzmann/pokedexcli/internal/commands"
	"os"
	"unicode"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var config = commands.Config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		command, ok := commands.GetCommands()[cleanedInput[0]]

		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.Callback(&config, cleanedInput)
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
