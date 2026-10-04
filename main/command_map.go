package main

import (
	"fmt"
)

func commandMapf(cfg *config, cliArgs... string) error {
	locationResp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextUrl, cfg.cache)
	if err != nil {
		return err
	}

	cfg.previousUrl = locationResp.Previous
	cfg.nextUrl = locationResp.Next

	for _, r := range locationResp.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandMapb(cfg *config, cliArg... string) error {
	if cfg.previousUrl == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocationArea(cfg.previousUrl, cfg.cache)
	if err != nil {
		return err
	}

	cfg.previousUrl = locationResp.Previous
	cfg.nextUrl = locationResp.Next

	for _, r := range locationResp.Results {
		fmt.Println(r.Name)
	}
	return nil
}
