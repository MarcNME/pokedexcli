package commands

import (
	"encoding/json"
	"fmt"
	"github.com/marc-enzmann/pokedexcli/internal/pokeapi"
)

func commandMap(config *Config) error {
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

func commandMapB(config *Config) error {
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

func callToLocationAreaApi(config *Config, url string) error {
	body, err := pokeapi.CallPokeApi(url)

	var locationArea LocationArea
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

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
