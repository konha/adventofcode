package main

import "testing"

func Test_sumOfResults(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			1320,
		},
		{
			"Values",
			"input.txt",
			515974,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfResults(lines)
			if got != tt.want {
				t.Errorf("sumOfResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_focusingPower(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			145,
		},
		{
			"Values",
			"input.txt",
			265894,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := focusingPower(lines)
			if got != tt.want {
				t.Errorf("focusingPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
