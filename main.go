package main

import (
	"time"

	"github.com/MichaelBo1/repldex/internal/pokeapi"
)

func main() {
	clientTimeout := 5 * time.Second
	purgeInterval := 10 * time.Second
	pokeapiClient := pokeapi.NewClient(clientTimeout, purgeInterval)

	conf := &cliConfig{
		api:     pokeapiClient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	startRepl(conf)
}
