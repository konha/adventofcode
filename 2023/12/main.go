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

	fmt.Println("Sum of Arrangements:", sumOfArrangements(lines))
	fmt.Println("Sum of Unfolded Arrangements:", sumOfArrangementsUnfolded(lines))
}

func sumOfArrangements(lines []string) int {
	sum := 0

	rows := parse(lines)
	for _, row := range rows {
		num := numOfArrangements(row)
		sum += num
	}

	return sum
}

func sumOfArrangementsUnfolded(lines []string) int {
	sum := 0

	rows := parse(lines)
	for _, row := range rows {
		unfoldedRow := unfoldedRow(row)
		noa := numOfArrangements(unfoldedRow)
		sum += noa
	}

	return sum
}

func unfoldedRow(row row) row {
	conditions := strings.Repeat(row.conditions+"?", 5)
	// remove last ?
	conditions = conditions[:len(conditions)-1]

	// repeat the row.damagedGroups 5 times
	damagedGroups := []int{}
	for i := 0; i < 5; i++ {
		damagedGroups = append(damagedGroups, row.damagedGroups...)
	}

	row.conditions = conditions
	row.damagedGroups = damagedGroups

	return row
}

type row struct {
	conditions    string
	damagedGroups []int
}

func numOfArrangements(row row) int {
	var cache [][]int

	for i := 0; i < len(row.conditions); i++ {
		cache = append(cache, make([]int, len(row.conditions)+1))
		for j := 0; j < len(row.conditions)+1; j++ {
			cache[i][j] = -1
		}
	}

	return calcNumOfArrangements(0, 0, row, cache)
}

func calcNumOfArrangements(index int, groupIndex int, row row, cache [][]int) int {
	if index == len(row.conditions) {
		if groupIndex < len(row.damagedGroups) {
			return 0
		}
		return 1
	}

	if cache[index][groupIndex] != -1 {
		return cache[index][groupIndex]
	}

	if row.conditions[index] == '.' {
		num := calcNumOfArrangements(index+1, groupIndex, row, cache)
		cache[index][groupIndex] = num
		return num
	}

	num := 0
	if row.conditions[index] == '?' {
		num += calcNumOfArrangements(index+1, groupIndex, row, cache)
	}
	if groupIndex < len(row.damagedGroups) {
		count := 0
		for i := index; i < len(row.conditions); i++ {
			if count > row.damagedGroups[groupIndex] || row.conditions[i] == '.' || count == row.damagedGroups[groupIndex] && row.conditions[i] == '?' {
				break
			}
			count++
		}
		if count == row.damagedGroups[groupIndex] {
			newIndex := index + count
			if newIndex < len(row.conditions) && row.conditions[newIndex] != '#' {
				newIndex++
			}
			num += calcNumOfArrangements(newIndex, groupIndex+1, row, cache)
		}
	}
	cache[index][groupIndex] = num

	return num
}

func validArrangement(conditions string, row row) bool {
	groups := strings.Split(conditions, ".")
	grouptCounts := []int{}

	for _, group := range groups {
		if len(group) == 0 {
			continue
		}
		grouptCounts = append(grouptCounts, len(group))
	}

	if len(grouptCounts) != len(row.damagedGroups) {
		return false
	}

	for i, damagedGroup := range row.damagedGroups {
		if grouptCounts[i] != damagedGroup {
			return false
		}
	}

	return true
}

func parse(lines []string) []row {
	rows := make([]row, 0, len(lines))
	for _, line := range lines {
		row := row{}
		parts := strings.Split(line, " ")
		row.conditions = parts[0]
		groups := strings.Split(parts[1], ",")
		for _, group := range groups {
			g, _ := strconv.Atoi(group)
			row.damagedGroups = append(row.damagedGroups, g)
		}
		rows = append(rows, row)
	}
	return rows
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
