package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		}, {
			input:    "one",
			expected: []string{"one"},
		}, {
			input:    "",
			expected: []string{},
		}, {
			input:    "I Do NOt KnOw",
			expected: []string{"i", "do", "not", "know"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length doesn't match expected.")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word %d doesn't match: got %s not %s", i, word, expectedWord)
				continue
			}
		}
	}
}
