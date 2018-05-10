package main

import "testing"

func TestGreeting(t *testing.T) {

	testmap := []struct {
		input    string
		expected string
	}{
		{
			input:    "Gophers",
			expected: "Greeting, Gophers",
		},
	}

	for _, tm := range testmap {
		output := Greeting(tm.input).String()
		if output != tm.expected {
			t.Errorf("got: %s wanted: %s", output, tm.expected)
		}
	}
}
