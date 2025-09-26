# Pokedex CLI

A simple command-line tool that lets you explore Pokémon locations, list all Pokémon living there, and even try to catch them.  
Built with Go and powered by [PokeAPI](https://pokeapi.co/).

## Features

- Browse Pokémon location areas, 20 at a time (`map`, `mapb`)
- Explore a location and see what Pokémon live there (`explore <location>`)
- Try to catch Pokémon with a probability mechanic (`catch <pokemon>`)
- View captured Pokémon and inspect their details (`pokedex`, `inspect <name>`)
- In-memory caching with automatic cleanup to reduce API calls

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/haanhtuandev/pokedexcli.git
cd pokedexcli
go build -o pokedex
./pokedex
