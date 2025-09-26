package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cf *Config, args ...string) error
}

type LocationList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Config struct {
	Next     string
	Previous string
}

var pokeMap = map[string]Pokemon{}

var url string = "https://pokeapi.co/api/v2/location-area"

var commandMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Provides user manual",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Subsequently displays 20 location areas in the Pokemon world",
		callback:    commandLoc,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays 20 previous location areas in the Pokemon world",
		callback:    commandLocBack,
	},
	"explore": {
		name:        "explore",
		description: "Takes the name of a location area as an argument.\nExample usage: explore <area_name>",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Trying to catch a Pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "view descriptions of a captured Pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "view all captured Pokemons",
		callback:    commandPokedex,
	},
}

func commandExit(cf *Config, s ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cf *Config, s ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n\n")
	// for key, value := range myCopyMap {
	// 	fmt.Printf("%v: %v", key, value.description)
	// }
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exits the Pokedex")
	fmt.Println("map: Next Pokemon world locations")
	fmt.Println("mapb: Previous Pokemon world locations")
	fmt.Println("explore [location name]: Get a list of Pokemon habitating that area")
	fmt.Println("catch [pokemon name]: Attempt to capture a Pokemon")
	fmt.Println("pokedex: View all Pokemoms captured")
	fmt.Print("\n")
	return nil
}

func commandLoc(cf *Config, s ...string) error {
	if cf.Next == "" {
		fmt.Println("you're on the first page!")
	}
	current_locations, err := retrieveLocation(cf.Next)

	if err != nil {
		return fmt.Errorf("error retrieving location: %w", err)
	}

	for _, loc := range current_locations.Results {
		fmt.Println(loc.Name)
	}
	cf.Previous = current_locations.Previous
	cf.Next = current_locations.Next

	return nil
}

func commandLocBack(cf *Config, s ...string) error {
	if cf.Previous == "" {
		fmt.Println("you're on the first page!")
	}
	current_locations, err := retrieveLocation(cf.Previous)

	if err != nil {
		return fmt.Errorf("error retrieving location: %w", err)
	}

	for _, loc := range current_locations.Results {
		fmt.Println(loc.Name)
	}
	cf.Next = current_locations.Next
	cf.Previous = current_locations.Previous
	return nil
}

func commandExplore(cf *Config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("explore command requires a location name")
	}
	exploreURL := fmt.Sprintf("%s/%s", url, name[0])

	loc, err := retrievePokemonHome(exploreURL)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name[0])
	fmt.Printf("Found Pokémon:\n")

	for _, pokemon := range loc.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cf *Config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("catch command requires a location name")
	}
	catchURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name[0])
	poc, err := retrievePokemon(catchURL)
	if err != nil {
		return err
	}
	userChance := rand.Intn(200)
	fmt.Printf("Throwing a Pokeball at %v...\n", poc.Name)
	time.Sleep(3000 * time.Millisecond)
	captureAttempt(userChance, poc)
	return nil
}

func captureAttempt(userChance int, pokemon Pokemon) {
	if userChance < pokemon.BaseExperience {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return
	} else {
		fmt.Printf("%v added to inventory\n", pokemon.Name)
		pokeMap[pokemon.Name] = pokemon
		return
	}

}

func commandInspect(cf *Config, name ...string) error {
	poke_name := name[0]
	val, exists := pokeMap[poke_name]
	if exists {
		fmt.Printf("Name: %v\n", val.Name)
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
	} else {
		fmt.Printf("You have not captured %v\n", poke_name)
	}
	return nil
}

func commandPokedex(cf *Config, name ...string) error {
	fmt.Printf("Pokedex: \n")
	for _, val := range pokeMap {
		fmt.Printf("- %v\n", val.Name)
	}
	return nil
}
