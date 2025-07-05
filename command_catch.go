package main

import (
	"fmt"
)

func commandCatch(cfg *config, arg1 *string) error {

	catchResp, err := cfg.pokeapiClient.PokemonInfo(nil, arg1)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *arg1)
	fmt.Println(catchResp.BaseExperience)

	return nil
}
