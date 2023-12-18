package main

import "testing"

func Test_sumOfValues(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			114,
		},
		{
			"Values",
			"input.txt",
			1921197370,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfValues(lines)
			if got != tt.want {
				t.Errorf("sumOfValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfValuesBackwards(t *testing.T) {
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
			"Values",
			"input.txt",
			1124,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfValuesBackwards(lines)
			if got != tt.want {
				t.Errorf("sumOfValuesBackwards() = %v, want %v", got, tt.want)
			}
		})
	}
}
