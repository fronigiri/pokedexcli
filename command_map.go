package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args []string) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationRes.Next
	cfg.nextLocationsURL = locationRes.Previous
	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationRes.Next
	cfg.nextLocationsURL = locationRes.Previous
	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}
	return nil

}
