package main

import (
	"time"

	"github.com/MrEkwunife/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeApiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
