package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocations(pageURL *string, area *string) (RespShallowExlpore, error) {

	url := baseUrl + "/location-area/" + *area
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		exploreResp := RespShallowExlpore{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return RespShallowExlpore{}, err
		}

		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowExlpore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowExlpore{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowExlpore{}, err
	}

	exploreResp := RespShallowExlpore{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespShallowExlpore{}, err
	}

	c.cache.Add(url, dat)

	return exploreResp, nil

}
