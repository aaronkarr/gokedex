package main

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name to check")
	}

	name := args[0]
	value, exists := cfg.pokedex[name]
	if !exists {
		return errors.New("you have not caught that pokemon yet")
	}

	fmt.Println("Name:", value.Name)
	fmt.Println("Height:", value.Height)
	fmt.Println("Weight:", value.Weight)
	fmt.Println("Stats:")
	for _, stat := range value.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Type:")
	for _, typeInfo := range value.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
