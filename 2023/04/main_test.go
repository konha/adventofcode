package main

import "testing"

func Test_sumOfPoints(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			13,
		},
		{
			"Values",
			"input.txt",
			26346,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfPoints(lines)
			if got != tt.want {
				t.Errorf("sumOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfScratchCards(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			30,
		},
		{
			"Values",
			"input.txt",
			8467762,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfScratchCards(lines)
			if got != tt.want {
				t.Errorf("sumOfScratchCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
