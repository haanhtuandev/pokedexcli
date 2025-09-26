package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	pokecache "github.com/haanhtuandev/pokedexcli/internal"
)

func cleanInput(text string) []string {
	// Convert the string to lowercase
	lowerCaseText := strings.ToLower(text)
	// Trim leading and trailing whitespace
	trimmedText := strings.TrimSpace(lowerCaseText)
	// Split the string by spaces and return the result
	return strings.Fields(trimmedText)
}

func retrieveLocation(url string) (LocationList, error) {
	var location LocationList
	cache, _ := pokecache.NewCache(5000 * time.Millisecond)
	cached_data, exist := cache.Get(url)
	if exist {
		err := json.Unmarshal(cached_data, &location)
		if err != nil {
			return location, fmt.Errorf("error saving data to empty Location struct: %w", err)
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return location, fmt.Errorf("error creating request: %w", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return location, fmt.Errorf("error getting response: %w", err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return location, fmt.Errorf("error converting response body to byte slice: %w", err)
	}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, fmt.Errorf("error saving data to empty Location struct: %w", err)
	}
	cache.Add(url, data)
	return location, nil
}

func retrievePokemonHome(url string) (LocationArea, error) {
	var loc LocationArea
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error crafting request: %v", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error receiving respond: %w", err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading response: %w", err)
	}
	json.Unmarshal(data, &loc)
	return loc, nil
}
func retrievePokemon(url string) (Pokemon, error) {
	var poc Pokemon
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error crafting request: %v", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error receiving respond: %w", err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response: %w", err)
	}
	json.Unmarshal(data, &poc)
	fmt.Println(poc.Name)
	return poc, nil

}

func handleCommand(input string) []string {
	return strings.Fields(input)
}
