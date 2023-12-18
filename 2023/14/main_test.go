package main

import "testing"

func Test_totalLoad(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			136,
		},
		{
			"Values",
			"input.txt",
			111339,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := totalLoad(lines)
			if got != tt.want {
				t.Errorf("totalLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalLoadCycles(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			64,
		},
		{
			"Values",
			"input.txt",
			93736,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := totalLoadCycles(lines)
			if got != tt.want {
				t.Errorf("totalLoadCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}
