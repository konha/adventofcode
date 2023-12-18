package main

import "testing"

func Test_cubicMeters1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			62,
		},
		{
			"Values",
			"input.txt",
			36679,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := cubicMeters(parse1(lines))
			if got != tt.want {
				t.Errorf("cubicMeters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cubicMeters2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			952408144115,
		},
		{
			"Values",
			"input.txt",
			88007104020978,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := cubicMeters(parse2(lines))
			if got != tt.want {
				t.Errorf("cubicMeters() = %v, want %v", got, tt.want)
			}
		})
	}
}
