package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MichaelBo1/repldex/internal/pokeapi"
)

type cliConfig struct {
	api              pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokémon in the given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon and add it to your Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the data for a previously caught Pokémon",
			callback:    commandInspect,
		},
	}
}

func startRepl(conf *cliConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := getCommands()[cmdName]
		if !ok {
			fmt.Println("unknown command")
			continue
		}

		err := cmd.callback(conf, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	out := strings.ToLower(input)
	words := strings.Fields(out)
	return words
}
