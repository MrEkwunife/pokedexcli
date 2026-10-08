package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
