package commands

import (
	"encoding/json"
	"fmt"
	"github.com/marc-enzmann/pokedexcli/internal/pokeapi"
)

func commandExplore(_ *Config, parameters []string) error {
	if len(parameters) < 2 {
		fmt.Printf("Please specify an area\nUseage: explore <location>\n")
		return nil
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", parameters[1])
	body, err := pokeapi.CallPokeApi(url)
	if err != nil {
		return err
	}

	var locationArea locationArea
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

type locationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int           `json:"min_level"`
				MaxLevel        int           `json:"max_level"`
				ConditionValues []interface{} `json:"condition_values"`
				Chance          int           `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
