package main

import "testing"

func Test_multipleOfWaysToBeatRecord(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			288,
		},
		{
			"Values",
			"input.txt",
			3316275,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := multipleOfWaysToBeatRecord(lines)
			if got != tt.want {
				t.Errorf("multipleOfWaysToBeatRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waysToBeatRecordLargeRace(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			71503,
		},
		{
			"Values",
			"input.txt",
			27102791,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := waysToBeatRecordLargeRace(lines)
			if got != tt.want {
				t.Errorf("waysToBeatRecordLargeRace() = %v, want %v", got, tt.want)
			}
		})
	}
}
