package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	pokemonName := args[0]
	pokemonInfo, exists := cfg.pokedex[pokemonName]
	if !exists {
		return errors.New("you have not caught this pokemon")
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")

	for _, s := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")

	for _, t := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil

}
