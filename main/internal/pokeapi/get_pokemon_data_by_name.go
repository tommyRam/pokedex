package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tommyRam/pokedexcli/internal/pokecache"
)

func (c *Client)GetPokemonData(pokemonName string, cache pokecache.Cache) (Pokemon, error){
	url := baseUrl + "/pokemon/" + pokemonName

	if byteData, ok := cache.Get(url); ok {
		// fmt.Println("Getting data from cache")
		return getPokemonStructFromByte(&byteData)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	// fmt.Print(res)
	if err != nil || res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("The pokemon %v is not found", pokemonName)
	}

	defer res.Body.Close()

	data, e := io.ReadAll(res.Body)
	if e != nil {
		return Pokemon{}, e
	}

	added_to_cache, err := cache.Add(url, data)
	if !added_to_cache {
		return Pokemon{}, err
	}

	return getPokemonStructFromByte(&data)
}

func getPokemonStructFromByte(byteData *[]byte) (Pokemon, error) {
	response :=  Pokemon{}
	err := json.Unmarshal(*byteData, &response)
	if err != nil {
		return Pokemon{}, err
	}
	return response, nil
}