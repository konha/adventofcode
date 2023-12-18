package main

import "testing"

func Test_sumOfSteps(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			2,
		},
		{
			"ExampleValues2",
			"input_example_2.txt",
			6,
		},
		{
			"Values",
			"input.txt",
			21797,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfSteps(lines)
			if got != tt.want {
				t.Errorf("sumOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfStepsGhost(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues3",
			"input_example_3.txt",
			6,
		},
		{
			"Values",
			"input.txt",
			23977527174353,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfStepsGhost(lines)
			if got != tt.want {
				t.Errorf("sumOfStepsGhost() = %v, want %v", got, tt.want)
			}
		})
	}
}
