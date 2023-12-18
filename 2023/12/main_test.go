package main

import "testing"

func Test_sumOfArrangements(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			21,
		},
		{
			"Values",
			"input.txt",
			6949,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfArrangements(lines)
			if got != tt.want {
				t.Errorf("sumOfArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfArrangementsUnfolded(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			525152,
		},
		{
			"Values",
			"input.txt",
			51456609952403,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfArrangementsUnfolded(lines)
			if got != tt.want {
				t.Errorf("sumOfArrangementsUnfolded() = %v, want %v", got, tt.want)
			}
		})
	}
}
