package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tommyRam/pokedexcli/internal/pokeapi"
	"github.com/tommyRam/pokedexcli/internal/pokecache"
)


type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, cliArgs... string) error
}

type config struct {
	commands      map[string]cliCommand
	pokemons      map[string]pokeapi.Pokemon
	cache         pokecache.Cache
	pokeapiClient pokeapi.Client
	nextUrl       *string
	previousUrl   *string
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description:"Get the next page of locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Get all encouters pokemon from the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch pokemon using its name",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Show details of caught pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Show all caught pokemon name",
			callback: commandPokedex,
		},
	}
	return commands
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	
	commands := cfg.commands

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		currentCommand := words[0]

		val, exist := commands[currentCommand]
		if !exist {
			fmt.Printf("Unknown command\n")
			continue
		}
		
		cliArgs :=  words[1:]

		err := val.callback(cfg, cliArgs...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	inputLower := strings.ToLower(input)
	words := strings.Fields(inputLower)
	return words
}