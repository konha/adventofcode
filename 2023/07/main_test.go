package main

import "testing"

func Test_totalWinnings(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			6440,
		},
		{
			"Values",
			"input.txt",
			247823654,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := totalWinnings(lines, false)
			if got != tt.want {
				t.Errorf("totalWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalWinningsSpecialJoker(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			5905,
		},
		{
			"Values",
			"input.txt",
			245461700,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := totalWinnings(lines, true)
			if got != tt.want {
				t.Errorf("totalWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
