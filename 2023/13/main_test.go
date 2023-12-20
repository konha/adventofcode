package main

import "testing"

func Test_sumOfPatterns(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			405,
		},
		{
			"Values",
			"input.txt",
			30487,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfPatterns(lines)
			if got != tt.want {
				t.Errorf("sumOfPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfPatternsSmudge(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			400,
		},
		{
			"Values",
			"input.txt",
			31954,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfPatternsSmudge(lines)
			if got != tt.want {
				t.Errorf("sumOfPatternsSmudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
