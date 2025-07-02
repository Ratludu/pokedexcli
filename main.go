package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var mapping = make(map[string]cliCommand)

func main() {

	mapping = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help messsge",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	fmt.Println("Welcome to Pokedex, Type a command to continue...")
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := mapping[line]; ok {
			run := mapping[line].callback
			run()
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("Pokedex > ")
	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	fmt.Println("")

	for command, val := range mapping {
		fmt.Printf("%s: %s\n", command, val.description)
	}

	return nil
}
func cleanInput(text string) []string {

	trimmed := strings.TrimSpace(text)
	if len(trimmed) == 0 {
		return []string{}
	}

	parts := strings.Split(trimmed, " ")

	cleaned := make([]string, len(parts))
	for i, word := range parts {
		cleaned[i] = strings.ToLower(word)
	}

	return cleaned
}
