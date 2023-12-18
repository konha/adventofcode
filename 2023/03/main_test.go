package main

import "testing"

func Test_sumOfPartNumbers(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			4361,
		},
		{
			"Values",
			"input.txt",
			543867,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfPartNumbers(lines)
			if got != tt.want {
				t.Errorf("sumOfPartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfGearRatios(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			467835,
		},
		{
			"Values",
			"input.txt",
			79613331,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfGearRatios(lines)
			if got != tt.want {
				t.Errorf("sumOfGearRatios() = %v, want %v", got, tt.want)
			}
		})
	}
}
