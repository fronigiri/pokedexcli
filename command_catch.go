package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonInfo, err := cfg.pokeapiClient.PokemonStats(pokemonName)
	if err != nil {
		return err
	}

	baseExperience := float64(pokemonInfo.BaseExperience)
	slope := -70.0 / 260.0
	intercept := 85.0 - slope*40.0
	catchChance := slope*baseExperience + intercept
	if catchChance < 15.0 {
		catchChance = 15.0
	} else if catchChance > 85.0 {
		catchChance = 85.0
	}
	if float64(rand.Intn(100)) < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonInfo
		fmt.Println("you may now inspect it with the inspect command")
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
}
