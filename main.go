package main

import "github.com/MrEkwunife/pokedexcli/internal/pokeapi"

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeApiClient: pokeapi.NewClient(),
	}
	startRepl(cfg)
}
