package main

import "fmt"

func commandExplore(cfg *config, arg1 *string) error {

	exploreResp, err := cfg.pokeapiClient.ExploreLocations(nil, arg1)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *arg1)
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}

	return nil
}
