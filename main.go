package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mharkness1/pokedexcli/internal/pokeapi"
)

// Defines global config which trackes page URLs relative.
type config struct {
	NextURL       string
	PreviousURL   string
	CaughtPokemon map[string]pokeapi.PokemonResults
}

// Creates new reader, prints welcome message once, and instantiates config.
func main() {
	r := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Type 'help' for list of commands.")
	cfg := &config{
		CaughtPokemon: make(map[string]pokeapi.PokemonResults),
	}

	// Continuous for loop, prints 'Pokedex >' and listens for input from the reader.
	for {
		fmt.Print("Pokedex > ")
		r.Scan()
		input := cleanInput(r.Text())

		//TODO more sophisticated passing arguments in, at the moment just checks first and second inputs.
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		// On input checks if it is in the dictionary. If it is in then calls the command and logs the error (config is always passed in).
		if i, ok := commandLookup[commandName]; !ok {
			fmt.Println("Unknown Command")
		} else {
			err := i.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

// Takes the input and formats to lower space and takes the first word "EXIT for the love of god" results in 'exit' command.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
