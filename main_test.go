package main

import (
	"reflect"
	"testing"
)

func TestGetFrequency(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		want map[rune]int
	}{
		{
			name: "Normal case",
			data: []rune("hello"),
			want: map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1},
		},
		{
			name: "Empty input",
			data: []rune(""),
			want: map[rune]int{},
		},
		{
			name: "Single element",
			data: []rune("a"),
			want: map[rune]int{'a': 1},
		},
		{
			name: "Repeated elements",
			data: []rune("aaa"),
			want: map[rune]int{'a': 3},
		},
		{
			name: "Case sensitivity",
			data: []rune("Aa"),
			want: map[rune]int{'A': 1, 'a': 1},
		},
		{
			name: "Non-ASCII runes",
			data: []rune("こんにちは"),
			want: map[rune]int{'こ': 1, 'ん': 1, 'に': 1, 'ち': 1, 'は': 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCharFrequency(tt.data)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFrequency(%q) = %v, want %v", string(tt.data), got, tt.want)
			}
		})
	}
}

func TestGetBTree(t *testing.T) {
	// Initial nodes based on your provided frequency data
	ls := nodes{
		leaf{char: 'C', frequency: 32},
		leaf{char: 'D', frequency: 42},
		leaf{char: 'E', frequency: 120},
		leaf{char: 'K', frequency: 7},
		leaf{char: 'L', frequency: 42},
		leaf{char: 'M', frequency: 24},
		leaf{char: 'U', frequency: 37},
		leaf{char: 'Z', frequency: 2},
	}

	// Build the tree
	root := getBTree(ls)

	// Check if the tree has been built correctly (example check)
	if root.weight() != 306 {
		t.Errorf("Expected weight of the root to be 306, got %d", root.weight())
	}
}
