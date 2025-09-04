package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	name := args[0]

	fmt.Printf("Exploring %s...\n", name)

	pokemonList, err := cfg.pokeapiClient.PokemonList(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pe := range pokemonList.PokemonEncounters {
		fmt.Printf(" - %s\n", pe.Pokemon.Name)
	}
	return nil
}
