package commands

import (
	"encoding/json"
	"fmt"

	"github.com/marc-enzmann/pokedexcli/internal/model"
	"github.com/marc-enzmann/pokedexcli/internal/pokeapi"
)

func commandMap(config *model.Config, _ []string) error {
	var url string
	if config.Next != "" {
		url = config.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	err2 := callToLocationAreaApi(config, url)
	if err2 != nil {
		return err2
	}

	return nil
}

func commandMapB(config *model.Config, _ []string) error {
	var url string
	if config.Previous != "" {
		url = config.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	err2 := callToLocationAreaApi(config, url)
	if err2 != nil {
		return err2
	}

	return nil
}

func callToLocationAreaApi(config *model.Config, url string) error {
	body, err := pokeapi.CallPokeApi(url)
	if err != nil {
		return err
	}

	var locationArea model.LocationAreas
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return err
	}

	config.Next = locationArea.Next
	config.Previous = locationArea.Previous

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}
