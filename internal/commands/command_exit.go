package commands

import (
	"fmt"
	"os"

	"github.com/marc-enzmann/pokedexcli/internal/model"
)

func commandExit(_ *model.Config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
