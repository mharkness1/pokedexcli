package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Defines global config which trackes page URLs relative.
type config struct {
	NextURL     string
	PreviousURL string
}

// Creates new reader, prints welcome message once, and instantiates config.
func main() {
	r := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Type 'help' for list of commands.")
	cfg := &config{}

	// Continuous for loop, prints 'Pokedex >' and listens for input from the reader.
	for {
		fmt.Print("Pokedex > ")
		r.Scan()
		input := cleanInput(r.Text())

		// On input checks if it is in the dictionary. If it is in then calls the command and logs the error (config is always passed in).
		if i, ok := commandLookup[input[0]]; !ok {
			fmt.Println("Unknown Command")
		} else {
			err := i.callback(cfg)
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
