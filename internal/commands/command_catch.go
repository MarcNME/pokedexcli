package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/marc-enzmann/pokedexcli/internal/model"
	"github.com/marc-enzmann/pokedexcli/internal/pokeapi"
)

func commandCatch(config *model.Config, args []string) error {
	if len(args) != 2 {
		return errors.New("you must provide a pokemon name\n")
	}

	var url = fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", args[1])
	pokemon, err := callToPokemonApi(url)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	config.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}

func callToPokemonApi(url string) (model.Pokemon, error) {
	body, err := pokeapi.CallPokeApi(url)
	if err != nil {
		return model.Pokemon{}, err
	}

	var pokemon model.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return model.Pokemon{}, err
	}

	return pokemon, nil
}
