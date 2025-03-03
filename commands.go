package main

import (
	"fmt"
	"os"

	"github.com/mharkness1/pokedexcli/internal/pokeapi"
)

// Properties of a command.
type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

// Create empty map taking command and mapping to the cliCommand struct (with callback)
var commandLookup map[string]cliCommand

// Exit
func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Lists the map of available commands and their descriptions.
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

// Calls for location areas from pokeapi and lists results (20 at a time), if active when called will move to next page.
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

// Will return to previous page of location areas when called. If called initally will return first 20.
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

// Init runs first as reserved function name, instatiates the map of commands and command words.
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
