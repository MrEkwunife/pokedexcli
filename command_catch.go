package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("usage: catch <pokeman-name>")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!\n", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
