package commands

import (
	"fmt"

	"github.com/marc-enzmann/pokedexcli/internal/model"
)

func commandHelp(_ *model.Config, _ []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range GetCommandsSorted() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}
