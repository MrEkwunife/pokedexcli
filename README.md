# pokedexcli

A lightweight, interactive command-line Pokédex REPL built with Go. Perfect for exploring the Pokémon world, scouting wild Pokémon across different regions, and catching them right from your terminal whenever you are bored or looking for a fun diversion.

---

## Motivation

This project was built as a hands-on way to practice real-world Go skills: building an interactive REPL, consuming a RESTful API, parsing JSON into structs, managing state, and implementing an in-memory cache with automatic expiry. It's also just a fun way to explore the Pokémon world from the terminal.

---

## Features

- **Interactive REPL**: A responsive command-line interface.
- **World Exploration**: Traverse location areas in batches of 20 with forward and backward pagination.
- **Area Scouting**: Discover which wild Pokémon inhabit specific location areas.
- **Catch Mechanic**: Throw Pokéballs to catch Pokémon, where catch difficulty scales with the Pokémon's base experience.
- **Pokédex & Inspection**: Keep track of caught Pokémon and inspect their stats, types, height, and weight.
- **In-Memory Caching**: Built-in cache with automatic reap intervals to minimize redundant network calls to the API.

---

## Quick Start

**Prerequisite:** [Go](https://go.dev/dl/) (version 1.22 or newer recommended)

```bash
go install github.com/MrEkwunife/pokedexcli@latest
pokedexcli
```

Make sure your Go binary path (typically `~/go/bin`) is included in your system's `PATH`.

To build from source instead:

```bash
git clone https://github.com/MrEkwunife/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
./pokedexcli
```

---

## Usage

Launch the application by running `pokedexcli` (or `./pokedexcli` if you built from source). You will enter the interactive REPL prompt:

```text
Pokedex >
```

### Quick Walkthrough

1. **Scout Locations**:
```text
   Pokedex > map
```
   Lists 20 location areas. Running `map` again lists the next 20, while `mapb` pages backward.

2. **Explore an Area**:
```text
   Pokedex > explore canalave-city-area
```
   Reveals all wild Pokémon residing in that area.

3. **Catch a Pokémon**:
```text
   Pokedex > catch tentacool
```
   Toss a Pokéball! You may need a few tries depending on the Pokémon's level and base experience.

4. **Inspect Your Catch**:
```text
   Pokedex > inspect tentacool
```
   Shows detailed stats (HP, Attack, Defense, Speed, etc.), height, weight, and types of any caught Pokémon.

5. **View Your Pokédex**:
```text
   Pokedex > pokedex
```
   Lists all Pokémon you have successfully caught.

### Available Commands

| Command | Arguments | Description |
|---|---|---|
| `help` | — | Displays available commands and usage instructions |
| `map` | — | Displays the next 20 location areas |
| `mapb` | — | Displays the previous 20 location areas |
| `explore` | `<location_area>` | Lists all Pokémon found in a specific location area |
| `catch` | `<pokemon_name>` | Attempts to catch a Pokémon and add it to your Pokédex |
| `inspect` | `<pokemon_name>` | Displays details and stats of a caught Pokémon |
| `pokedex` | — | Displays all Pokémon currently in your Pokédex |
| `exit` | — | Exits the Pokédex REPL |

---

## Contributing

Contributions are welcome! To get started:

1. Fork the repository on GitHub.
2. Clone your fork:
```bash
   git clone https://github.com/<your-username>/pokedexcli.git
   cd pokedexcli
```
3. Create a feature branch:
```bash
   git checkout -b my-feature
```
4. Make your changes and run the tests:
```bash
   go test ./...
```
5. Commit, push to your fork, and open a pull request against `main` describing what you changed and why.

Bug reports and feature ideas are also welcome via GitHub Issues.

---

## Attributions

- [PokéAPI](https://pokeapi.co/) — for providing the free RESTful Pokémon API.
- [JSON-to-Go](https://mholt.github.io/json-to-go/) by Matt Holt — for converting PokéAPI JSON responses into Go struct definitions.
