package pokedexapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonList(name string) (RespPokemonEncounters, error) {
	url := baseURL + "/location-area/" + name + "/"
	pokemonResp := RespPokemonEncounters{}
	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return RespPokemonEncounters{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	defer resp.Body.Close()

	dat2, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	err = json.Unmarshal(dat2, &pokemonResp)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	c.cache.Add(url, dat2)
	return pokemonResp, nil
}
