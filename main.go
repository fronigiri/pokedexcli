package main

import (
	"pokedexcli/internal/pokedexapi"
	"time"
)

func main() {
	pokeClient := pokedexapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
