package pokedexapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	locationsResp := RespShallowLocations{}
	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat2, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(dat2, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, dat2)
	return locationsResp, nil
}
