# PokeAPI REPL

This project is a simple Read-Eval-Print Loop (REPL) written in Go, designed to interact with the [PokeAPI](https://pokeapi.co/). The REPL allows you to explore Pokemon data and build up a personal Pokedex interactively.

## Purpose
The primary goal of this project was to learn the basics of Go programming, including making HTTP requests, handling JSON, and building a simple REPL interface.

## Features
- Search for Pokemon by name or ID to retrieve details like types, abilities, and stats.
- Add Pokemon to your Pokedex for personal reference.
- View your Pokedex to see the list of Pokemon you've added.
- Explore different locations in the Pokemon games.
- Simple and interactive command-line interface.

## Requirements
- Go 1.23+
- Internet connection (to fetch data from PokeAPI)

## Getting Started

### Clone the Repository

`git clone https://github.com/your-username/pokeapi-repl.git](https://github.com/MichaelBo1/REPLdex.git`

### Run the REPL

`go run main.go`

### Usage
1. Start the REPL by running the application.
2. Use the following commands to interact:
Use the following commands to interact:
- `help`: Displays a help message.
- `exit`: Exit the Pokedex.
- `map`: List locations.
- `mapb`: List previous locations.
- `explore <area>`: List all Pokémon in the given area.
- `catch <pokemon-name-or-id>`: Catch a Pokemon and add it to your Pokedex.
- `inspect <pokemon-name-or-id>`: Inspect the data for a previously caught Pokémon.
- `pokedex`: List all caught Pokemon.

## Example Interaction
```
> explore forest
Pokemon found in forest:
1. Caterpie
2. Weedle

> catch caterpie
Caterpie added to your Pokedex!

> pokedex
Your Pokedex:
1. Caterpie (Bug)

> inspect caterpie
Name: Caterpie
Type: Bug
Abilities: Shield Dust, Run Away
Stats: HP: 45, Attack: 30, Defense: 35, ...

> exit
```

## Learning Goals
This project covers several Go programming concepts:
- Basic syntax and structure of Go.
- Making HTTP requests and working with APIs.
- Parsing and handling JSON data.
- Building a simple REPL interface.
- Managing dependencies and structuring Go projects.

## Acknowledgements
- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data.

