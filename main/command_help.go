package main

import "fmt"

func commandHelp(cfg *config, cliArgs... string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cli := range cfg.commands {
		fmt.Printf("%s: %s\n", cli.name, cli.description)
	}

	return nil
}