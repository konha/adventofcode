package main

import "testing"

func Test_sumOfCalibrationValues(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		f        func(string) int
		want     int
	}{
		{
			"ExampleValues",
			"input_example.txt",
			calibrationValue,
			142,
		},
		{
			"Values",
			"input.txt",
			calibrationValue,
			55029,
		},
		{
			"ExampleValuesWithWords",
			"input_example_2.txt",
			calibrationValueWithWords,
			281,
		},
		{
			"ValuesWithWords",
			"input.txt",
			calibrationValueWithWords,
			55686,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := readFile(tt.filename)
			if err != nil {
				t.Errorf("readFile() error = %v", err)
			}
			got := sumOfCalibrationValues(lines, tt.f)
			if got != tt.want {
				t.Errorf("sumOfCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
