package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(conf *cliConfig) error {
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

func commandExit(conf *cliConfig) error {
	os.Exit(0)
	return nil
}

func commandMap(conf *cliConfig) error {
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

func commandMapB(conf *cliConfig) error {
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
