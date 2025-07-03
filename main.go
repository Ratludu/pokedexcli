package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
		"map": {
			name:        "map",
			description: "LocationArea",
			callback:    getLocationAreas,
		},
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

type locationArea struct {
	Count int `json:"count"`
}

func getLocationAreas() error {

	url := "https://pokeapi.co/api/v2/location-area/"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Recieved non-Ok HTTP Status Code: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var locationarea locationArea

	err = json.Unmarshal(body, &locationarea)
	if err != nil {
		log.Fatalf("Error unmarshalling the body: %v", err)
	}

	fmt.Println(locationarea.Count)

	return nil

}
