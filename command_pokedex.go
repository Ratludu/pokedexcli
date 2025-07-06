package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg1 *string) error {

	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex it empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("	- %s\n", pokemon.Name)
	}

	return nil
}
