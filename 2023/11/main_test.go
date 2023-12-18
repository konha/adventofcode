package main

import "testing"

func Test_sumOfShortestPaths(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			374,
		},
		{
			"Values",
			"input.txt",
			9742154,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfShortestPaths(lines)
			if got != tt.want {
				t.Errorf("thing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfShortestPathsHuge(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"Values",
			"input.txt",
			411142919886,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfShortestPathsHuge(lines)
			if got != tt.want {
				t.Errorf("sumOfShortestPathsHuge() = %v, want %v", got, tt.want)
			}
		})
	}
}
