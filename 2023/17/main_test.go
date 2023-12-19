package main

import "testing"

func Test_leastHeatLoss(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			102,
		},
		{
			"Values",
			"input.txt",
			817,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := leastHeatLoss(lines)
			if got != tt.want {
				t.Errorf("leastHeatLoss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leastHeatLoss2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			94,
		},
		{
			"ExampleValues2",
			"input_example_2.txt",
			71,
		},
		{
			"Values",
			"input.txt",
			925,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := leastHeatLoss2(lines)
			if got != tt.want {
				t.Errorf("leastHeatLoss2() = %v, want %v", got, tt.want)
			}
		})
	}
}
