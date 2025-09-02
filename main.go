package main

import (
	"pokedexcli/internal/pokedexapi"
	"time"
)

func main() {
	pokeClient := pokedexapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
