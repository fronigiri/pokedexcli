package pokedexapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonStats(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name + "/"
	pokemonResp := RespPokemon{}
	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat2, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	err = json.Unmarshal(dat2, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}
	c.cache.Add(url, dat2)
	return pokemonResp, nil
}
