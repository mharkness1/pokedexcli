package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	NextURL     string
	PreviousURL string
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Type 'help' for list of commands.")
	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")
		r.Scan()
		input := cleanInput(r.Text())

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

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
