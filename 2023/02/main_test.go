package main

import "testing"

func Test_sumOfValidGameIDs(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		max_red   int
		max_green int
		max_blue  int
		want      int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			12,
			13,
			14,
			8,
		},
		{
			"Values",
			"input.txt",
			12,
			13,
			14,
			2204,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfValidGameIDs(lines, tt.max_red, tt.max_green, tt.max_blue)
			if got != tt.want {
				t.Errorf("sumOfValidGameIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfPowerOfMinSets(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			2286,
		},
		{
			"Values",
			"input.txt",
			71036,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfPowerOfMinSets(lines)
			if got != tt.want {
				t.Errorf("sumOfPowerOfMinSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
