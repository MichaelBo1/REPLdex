package main

import (
	"time"

	"github.com/MichaelBo1/repldex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	conf := &cliConfig{
		api: pokeapiClient,
	}

	startRepl(conf)
}
