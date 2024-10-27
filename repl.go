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
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
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

		cmd, ok := getCommands()[cmdName]
		if !ok {
			fmt.Println("unknown command")
			continue
		}

		err := cmd.callback(conf)
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
