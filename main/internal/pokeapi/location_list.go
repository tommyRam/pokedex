package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tommyRam/pokedexcli/internal/pokecache"
)

func (c *Client)ListLocationArea(pageUrl *string, cache pokecache.Cache) (RespShallowLocations, error){
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if byteData, ok := cache.Get(url); ok {
		fmt.Println("Getting data from cache")
		return getLocationStructFromByte(&byteData)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()

	data, e := io.ReadAll(res.Body)
	if e != nil {
		return RespShallowLocations{}, e
	}

	added_to_cache, err := cache.Add(url, data)
	if !added_to_cache {
		return RespShallowLocations{}, err
	}

	return getLocationStructFromByte(&data)
}

func getLocationStructFromByte(byteData *[]byte) (RespShallowLocations, error) {
	response :=  RespShallowLocations{}
	err := json.Unmarshal(*byteData, &response)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return response, nil
}