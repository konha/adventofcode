package main

import "testing"

func Test_stepsToFarthestPoint(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			4,
		},
		{
			"ExampleValues",
			"input_example_2.txt",
			8,
		},
		{
			"Values",
			"input.txt",
			6820,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := stepsToFarthestPoint(lines)
			if got != tt.want {
				t.Errorf("stepsToFarthestPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_enclosedTiles(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example_3.txt",
			4,
		},
		{
			"ExampleValues",
			"input_example_4.txt",
			8,
		},
		{
			"ExampleValues",
			"input_example_5.txt",
			10,
		},
		{
			"Values",
			"input.txt",
			337,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := enclosedTiles(lines)
			if got != tt.want {
				t.Errorf("enclosedTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
