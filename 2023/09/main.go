package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("filename is empty")
		return
	}
	filename := os.Args[1]

	lines, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sum of Values:", sumOfValues(lines))
	fmt.Println("Sum of Values (Backwards):", sumOfValuesBackwards(lines))
}

func sumOfValues(lines []string) int {
	values := parse(lines)

	sum := 0
	for _, v := range values {
		sum += nextValue(v)
	}

	return sum
}

func sumOfValuesBackwards(lines []string) int {
	values := parse(lines)

	sum := 0
	for _, v := range values {
		sum += nextValueBackwards(v)
	}

	return sum
}

func nextValue(values []int) int {
	diffs := make([][]int, 0)

	diffs = append(diffs, values)
	diff := differences(values)
	diffs = append(diffs, diff)

	for !allZeroes(diff) {
		diff = differences(diff)
		diffs = append(diffs, diff)
	}

	nextValue := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		// last item of line
		lastValue := diffs[i][len(diffs[i])-1]

		if i == 0 {
			nextValue = lastValue
			continue
		}

		previousLine := diffs[i-1]
		previousLastValue := previousLine[len(previousLine)-1]
		newValue := lastValue + previousLastValue
		newLine := make([]int, len(previousLine)+1)
		copy(newLine, previousLine)
		newLine[len(newLine)-1] = newValue
		diffs[i-1] = newLine
	}

	return nextValue
}

func nextValueBackwards(values []int) int {
	diffs := make([][]int, 0)

	diffs = append(diffs, values)
	diff := differences(values)
	diffs = append(diffs, diff)

	for !allZeroes(diff) {
		diff = differences(diff)
		diffs = append(diffs, diff)
	}

	nextValue := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		firstValue := diffs[i][0]
		if i == 0 {
			nextValue = firstValue
			continue
		}

		previousLine := diffs[i-1]
		previousFirstValue := previousLine[0]

		newValue := previousFirstValue - firstValue
		newLine := make([]int, len(previousLine)+1)
		copy(newLine[1:], previousLine)
		newLine[0] = newValue
		diffs[i-1] = newLine
	}

	return nextValue
}

func differences(values []int) []int {
	diffs := make([]int, len(values)-1)

	for i := 1; i < len(values); i++ {
		diffs[i-1] = values[i] - values[i-1]
	}

	return diffs
}

func allZeroes(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func parse(lines []string) [][]int {
	values := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		values[i] = make([]int, len(fields))
		for j, field := range fields {
			val, _ := strconv.Atoi(field)
			values[i][j] = val
		}
	}

	return values
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
