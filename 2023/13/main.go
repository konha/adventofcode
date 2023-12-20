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
	filename := os.Args[1]

	lines, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sum of Patterns Part 1:", sumOfPatterns(lines))
	fmt.Println("Sum of Patterns Part 2:", sumOfPatternsSmudge(lines))
}

type pattern []string
type patterns []pattern

func sumOfPatterns(lines []string) int {
	patterns := parse(lines)

	cols := 0
	rows := 0

	for _, p := range patterns {
		c := colsBeforeVerticalReflection(p, 0)
		cols += c
		r := rowsAboveHorizontalReflection(p, 0)
		rows += r
	}

	return cols + 100*rows
}

func rowsAboveHorizontalReflection(p pattern, compare int) int {
	rows := 0

	for i := 1; i < len(p); i++ {
		if isHorizontalReflection(p, i, i-1) && i != compare {
			rows = i
			break
		}
	}

	return rows
}

func colsBeforeVerticalReflection(p pattern, compare int) int {
	cols := 0

	for i := 1; i < len(p[0]); i++ {
		if isVerticalReflection(p, i, i-1) && i != compare {
			cols = i
			break
		}
	}

	return cols
}

func isHorizontalReflection(p pattern, i, j int) bool {
	if i > len(p)-1 || j < 0 {
		return true
	}
	return p[i] == p[j] && isHorizontalReflection(p, i+1, j-1)
}

func isVerticalReflection(p pattern, i, j int) bool {
	if i > len(p[0])-1 || j < 0 {
		return true
	}
	identical := true
	for _, row := range p {
		if row[i] != row[j] {
			identical = false
			break
		}
	}

	return identical && isVerticalReflection(p, i+1, j-1)
}

func sumOfPatternsSmudge(lines []string) int {
	result := 0
	patterns := parse(lines)

	for _, p := range patterns {
		c := colsBeforeVerticalReflection(p, 0)
		r := rowsAboveHorizontalReflection(p, 0)
		result += patternSmudge(p, 0, 0, c, r)
	}

	return result
}

func patternSmudge(p pattern, i, j, cmpCols, cmpRows int) int {
	if i == len(p) || j == len(p[i]) {
		return 0
	}

	// change the char at i,j)
	flipChar(p, i, j)

	e := colsBeforeVerticalReflection(p, cmpCols) + 100*rowsAboveHorizontalReflection(p, cmpRows)
	if e != 0 {
		return e
	}

	// change back
	flipChar(p, i, j)

	if j == len(p[i])-1 {
		j = 0
		i = i + 1
	} else {
		j++
	}

	return patternSmudge(p, i, j, cmpCols, cmpRows)
}

func flipChar(input []string, i, j int) {
	line := strings.Split(input[i], "")
	if line[j] == "#" {
		line[j] = "."
	} else {
		line[j] = "#"
	}
	input[i] = strings.Join(line, "")
}

func parse(lines []string) patterns {
	ps := make(patterns, 0)

	p := make(pattern, 0)
	for i, line := range lines {
		if line == "" {
			ps = append(ps, p)
			p = make(pattern, 0)
			continue
		}
		p = append(p, line)
		if i == len(lines)-1 {
			ps = append(ps, p)
		}
	}

	return ps
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
