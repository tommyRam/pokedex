package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, cliArgs... string) error {
	if len(cliArgs) != 1 {
		return fmt.Errorf("No pokemon name was provided. eg. catch <pokemon_name>")
	}

	pokemonName := cliArgs[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	_, ok := cfg.pokemons[pokemonName]
	if ok {
		fmt.Printf("%v was caught!\n", pokemonName) 
	}

	pokemonData, err := cfg.pokeapiClient.GetPokemonData(pokemonName, cfg.cache)
	if err != nil {
		return err
	}

	currentChanceToCatch := rand.Intn(pokemonData.BaseExperience)
	if currentChanceToCatch >= (pokemonData.BaseExperience / 2) {
		cfg.pokemons[pokemonName] = pokemonData
		fmt.Printf("%v was caught!\n", pokemonName) 
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}
	return nil
}