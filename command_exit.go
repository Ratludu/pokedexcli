package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, arg1 *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
