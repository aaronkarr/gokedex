package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cliCommand := range getCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	fmt.Println()
	return nil
}
