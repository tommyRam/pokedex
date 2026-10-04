package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tommyRam/pokedexcli/internal/pokecache"
)

func (c *Client) List_pokemon_in_location(locationArea string, cache pokecache.Cache) (RespPokemonsInLocation, error) {
	url := baseUrl + "/location-area/" + locationArea

	if byteData, cached_data := cache.Get(url); cached_data {
		data, err := getPokemonsInLocationFromByte(byteData)
		return data, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonsInLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil || res.StatusCode > 399 {
		return RespPokemonsInLocation{}, fmt.Errorf("Can't fetch the data")
	}

	defer res.Body.Close()

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonsInLocation{}, err
	}

	cache.Add(url, byteData)
	return getPokemonsInLocationFromByte(byteData)
}

func getPokemonsInLocationFromByte(byteData []byte) (RespPokemonsInLocation, error) {
	var respPokemonsInLocation RespPokemonsInLocation
	err := json.Unmarshal(byteData, &respPokemonsInLocation)
	if err != nil {
		return RespPokemonsInLocation{}, nil
	}
	return respPokemonsInLocation, nil
}