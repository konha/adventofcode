package main

import "testing"

func Test_sumOfRatingNumbersOfAcceptedParts(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			19114,
		},
		{
			"Values",
			"input.txt",
			383682,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfRatingNumbersOfAcceptedParts(lines)
			if got != tt.want {
				t.Errorf("sumOfRatingNumbersOfAcceptedParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distinctCombinationsOfRatings(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			167409079868000,
		},
		{
			"Values",
			"input.txt",
			117954800808317,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := distinctCombinationsOfRatings(lines)
			if got != tt.want {
				t.Errorf("distinctCombinationsOfRatings() = %v, want %v", got, tt.want)
			}
		})
	}
}
