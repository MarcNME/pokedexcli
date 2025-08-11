package commands

import (
	"fmt"

	"github.com/marc-enzmann/pokedexcli/internal/model"
)

func commandPokedex(c *model.Config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.CaughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
