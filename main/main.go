package main

import (
	"time"

	"github.com/tommyRam/pokedexcli/internal/pokeapi"
	"github.com/tommyRam/pokedexcli/internal/pokecache"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)

	cfg := config{
		commands: getCommands(),
		pokeapiClient: client,
		cache: cache,
		pokemons: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}