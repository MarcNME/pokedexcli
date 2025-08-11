package commands

import (
	"encoding/json"
	"fmt"

	"github.com/marc-enzmann/pokedexcli/internal/model"
	"github.com/marc-enzmann/pokedexcli/internal/pokeapi"
)

func commandExplore(_ *model.Config, parameters []string) error {
	if len(parameters) < 2 {
		fmt.Printf("Please specify an area\nUseage: explore <location>\n")
		return nil
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", parameters[1])
	body, err := pokeapi.CallPokeApi(url)
	if err != nil {
		return err
	}

	var locationArea model.LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Explorting %s\n", locationArea.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println("  - " + encounter.Pokemon.Name)
	}

	return nil
}
