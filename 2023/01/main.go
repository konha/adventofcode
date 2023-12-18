package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("filename is empty")
		return
	}
	lines, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sum of calibration values:", sumOfCalibrationValues(lines, calibrationValue))
	fmt.Println("Sum of calibration values with words:", sumOfCalibrationValues(lines, calibrationValueWithWords))
}

func sumOfCalibrationValues(lines []string, f func(string) int) int {
	sum := 0

	for _, line := range lines {
		calibrationValue := f(line)
		sum += calibrationValue
	}

	return sum
}

func calibrationValue(line string) int {
	v := 0
	first := ""
	last := ""

	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			if first == "" {
				first = string(line[i])
			}
			last = string(line[i])
		}
	}

	fmt.Sscanf(first+last, "%d", &v)

	return v
}

func calibrationValueWithWords(line string) int {
	digits := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var result strings.Builder
	i := 0
	for i < len(line) {
		found := false
		for word, digit := range digits {
			if strings.HasPrefix(line[i:], word) {
				result.WriteString(digit)
				i++
				found = true
				break
			}
		}
		if !found {
			result.WriteByte(line[i])
			i++
		}
	}

	return calibrationValue(result.String())
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
