package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}

	if dat, ok := c.cache.Get(fullUrl); ok {
		fmt.Println("cache hit!")
		locationAreaResp := LocationAreasResp{}
		if err := json.Unmarshal(dat, &locationAreaResp); err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreaResp, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationArea := LocationAreasResp{}
	if err := json.Unmarshal(dat, &locationArea); err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullUrl, dat)
	return locationArea, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	if dat, ok := c.cache.Get(fullUrl); ok {
		fmt.Println("cache hit!")
		locationArea := LocationArea{}
		if err := json.Unmarshal(dat, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	if err := json.Unmarshal(dat, &locationArea); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, dat)
	return locationArea, nil
}
