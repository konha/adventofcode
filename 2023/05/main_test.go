package main

import "testing"

func Test_lowestLocationNum(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			35,
		},
		{
			"Values",
			"input.txt",
			806029445,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := lowestLocationNum(lines)
			if got != tt.want {
				t.Errorf("lowestLocationNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestLocationNum3(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			46,
		},
		{
			"Values",
			"input.txt",
			59370572,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := lowestLocationNum2(lines)
			if got != tt.want {
				t.Errorf("lowestLocationNum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
