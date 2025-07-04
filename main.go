package main

import (
	"github.com/ratludu/pokedexcli/internal/pokeapi"
	"time"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
