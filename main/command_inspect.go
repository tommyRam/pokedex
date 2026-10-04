package main

import (
	"fmt"
)

func commandInspect(cfg *config, cliArgs... string) error {
	if len(cliArgs) != 1 {
		return fmt.Errorf("Pokemon name not provided. eg. inspect <pokemon_name>")
	}

	pokemonName := cliArgs[0]
	data, ok := cfg.pokemons[pokemonName]
	if !ok {
		return fmt.Errorf("Pokemon %v not caught yet!", pokemonName)
	}

	fmt.Printf("Name: %v\n", pokemonName)
	fmt.Printf("Weight: %v\n", data.Weight)
	fmt.Printf("Stats:\n")

	for _, stat := range data.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range data.Types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}
	return nil
}