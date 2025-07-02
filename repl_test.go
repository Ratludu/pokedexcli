package main

import (
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Clean text: got %d want %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := strings.ToLower(actual[i])
			expectedWord := strings.ToLower(c.expected[i])

			if word != expectedWord {
				t.Errorf("Match Error: got: %s want: %s", word, expectedWord)
			}

		}
	}
}
