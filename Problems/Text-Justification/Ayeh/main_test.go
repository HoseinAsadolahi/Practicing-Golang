package main

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleFullJustify() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	x := fullJustify(words, maxWidth)
	for _, line := range x {
		fmt.Println(line)
	}
	//Output:
	//This    is    an
	//example  of text
	//justification.
}

func BenchmarkName(b *testing.B) {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 16
	for i := 0; i < b.N; i++ {
		fullJustify(words, maxWidth)
	}
}

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		maxWidth int
		expected []string
	}{
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.words, "_"), func(t *testing.T) {
			actual := fullJustify(test.words, test.maxWidth)
			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d lines, but got %d", len(test.expected), len(actual))
			}

			for i := 0; i < len(actual); i++ {
				if actual[i] != test.expected[i] {
					t.Errorf("Line %d: Expected '%s', but got '%s'", i, test.expected[i], actual[i])
				}
			}
		})
	}
}
