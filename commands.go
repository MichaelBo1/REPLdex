package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(conf *cliConfig, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(conf *cliConfig, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(conf *cliConfig, args ...string) error {
	res, err := conf.api.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = res.Next
	conf.prevLocationsURL = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(conf *cliConfig, args ...string) error {
	if conf.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	res, err := conf.api.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = res.Next
	conf.prevLocationsURL = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(conf *cliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("location name must be provided")
	}
	location := args[0]

	res, err := conf.api.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Pokemon:")
	for _, enc := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
