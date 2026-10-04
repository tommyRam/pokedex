package main

import "fmt"

func commandExplore(cfg *config, cliArgs... string) error {
	if len(cliArgs) != 1 {
		return fmt.Errorf("No argument was provided. Please do eg. explore <name_of_location_area>")
	}

	pokemons, err := cfg.pokeapiClient.List_pokemon_in_location(cliArgs[0], cfg.cache)
	if err != nil {
		return err
	}

	// fmt.Printf("Res: %v\n", pokemons.PokemonEncounters)
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}