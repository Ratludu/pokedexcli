package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg1 *string) error {

	const maxValue int = 500

	catchResp, err := cfg.pokeapiClient.PokemonInfo(nil, arg1)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *arg1)

	randNum := rand.Intn(maxValue)
	if randNum > catchResp.BaseExperience {
		fmt.Printf("%s was caught!\n", *arg1)
		cfg.pokedex[*arg1] = catchResp
	} else {
		fmt.Printf("%s escaped!\n", *arg1)
	}

	return nil
}
