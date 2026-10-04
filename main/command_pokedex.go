package main

import (
	"fmt"
)

func commandPokedex(cfg *config, cliArgs... string) error {
	fmt.Println("Your pokedex:")
	for k, _ := range cfg.pokemons {
		fmt.Printf(" - %v\n", k)
	}
	return nil
}