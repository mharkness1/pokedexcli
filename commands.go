package main

import (
	"fmt"
	"os"

	"github.com/mharkness1/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

var commandLookup map[string]cliCommand

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for key, val := range commandLookup {
		fmt.Printf("%s: %v\n", key, val.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(config *config) error {
	client := pokeapi.NewClient()

	locations, err := client.GetLocationAreas(config.NextURL)
	if err != nil {
		return err
	}

	config.NextURL = locations.Next
	config.PreviousURL = locations.Previous

	for i := range locations.Results {
		fmt.Println(locations.Results[i].Name)
	}
	return nil
}

func commandMapb(config *config) error {
	client := pokeapi.NewClient()

	locations, err := client.GetLocationAreas(config.PreviousURL)
	if err != nil {
		return err
	}

	config.NextURL = locations.Next
	config.PreviousURL = locations.Previous

	for i := range locations.Results {
		fmt.Println(locations.Results[i].Name)
	}
	return nil
}

func init() {
	commandLookup = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next/first page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of locations",
			callback:    commandMapb,
		},
	}
}
