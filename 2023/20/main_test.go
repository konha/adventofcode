package main

import "testing"

func Test_pulses(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			32000000,
		},
		{
			"ExampleValues2",
			"input_example_2.txt",
			11687500,
		},
		{
			"Values",
			"input.txt",
			821985143,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := pulses(lines)
			if got != tt.want {
				t.Errorf("pulses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fewestButtonPresses(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"Values",
			"input.txt",
			240853834793347,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := fewestButtonPresses(lines)
			if got != tt.want {
				t.Errorf("fewestButtonPresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
