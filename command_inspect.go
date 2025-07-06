package main

import (
	"fmt"
)

func commandInspect(cfg *config, arg1 *string) error {

	pokemon, ok := cfg.pokedex[*arg1]
	if !ok {
		fmt.Println("This Pokemon has not been caught yet!")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, i := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", i.Stat.Name, i.BaseStat)
	}
	fmt.Println("Types:")

	for _, i := range pokemon.Types {
		fmt.Printf("	- %s\n", i.Type.Name)
	}

	return nil
}
