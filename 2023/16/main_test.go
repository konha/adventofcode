package main

import "testing"

func Test_countEnergizedTiles(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			46,
		},
		{
			"Values",
			"input.txt",
			7307,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := countEnergizedTiles(lines)
			if got != tt.want {
				t.Errorf("countEnergizedTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEnergizedTiles(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			51,
		},
		{
			"Values",
			"input.txt",
			7635,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := maxEnergizedTiles(lines)
			if got != tt.want {
				t.Errorf("countEnergizedTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
