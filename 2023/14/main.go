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

	fmt.Println("Total Load Part 1:", totalLoad(lines))
	fmt.Println("Total Load Part 2:", totalLoadCycles(lines))
}

func totalLoad(lines []string) int {
	tilted := tilt(lines)
	return calcLoad(tilted)
}

func totalLoadCycles(lines []string) int {
	platform := lines

	// the result after 1000 cycles is the same as after 1000000000 cycles
	for i := 0; i < 1000; i++ {
		platform = tiltCycle(platform)
	}

	return calcLoad(platform)
}

func calcLoad(lines []string) int {
	load := 0

	max := len(lines[0])
	for i, line := range lines {
		lineWeight := max - i
		for _, char := range line {
			if char == 'O' {
				load += lineWeight
			}
		}
	}

	return load
}

func tilt(lines []string) []string {
	platform := lines

	for i, line := range lines {
		if i == 0 {
			continue
		}
		for j, char := range line {
			if char == 'O' {
				// move up until we hit a '#', an 'O' or the top of the platform
				// put the 'O' on the last place that was a '.'
				moveTo := i
				for k := i - 1; k >= 0; k-- {
					if platform[k][j] == '#' || platform[k][j] == 'O' {
						// we store it on the line below
						moveTo = k + 1
						break
					}
					if k == 0 {
						moveTo = 0
					}
				}
				// put a '.' on the previous place (current line)
				platform[i] = platform[i][:j] + "." + platform[i][j+1:]
				// put the 'O' on the moveTo line
				platform[moveTo] = platform[moveTo][:j] + "O" + platform[moveTo][j+1:]
			}
		}
	}

	return platform
}

func tiltCycle(lines []string) []string {
	platform := lines

	// tilt north
	for i, line := range lines {
		if i == 0 {
			continue
		}
		for j, char := range line {
			if char == 'O' {
				moveTo := i
				for k := i - 1; k >= 0; k-- {
					if platform[k][j] == '#' || platform[k][j] == 'O' {
						moveTo = k + 1
						break
					}
					if k == 0 {
						moveTo = 0
					}
				}
				platform[i] = platform[i][:j] + "." + platform[i][j+1:]
				platform[moveTo] = platform[moveTo][:j] + "O" + platform[moveTo][j+1:]
			}
		}
	}

	// tilt west
	for i := 0; i < len(lines); i++ {
		for j, char := range lines[i] {
			if j == 0 {
				continue
			}
			if char == 'O' {
				moveTo := j
				for k := j - 1; k >= 0; k-- {
					if platform[i][k] == '#' || platform[i][k] == 'O' {
						moveTo = k + 1
						break
					}
					if k == 0 {
						moveTo = 0
					}
				}
				platform[i] = platform[i][:j] + "." + platform[i][j+1:]
				platform[i] = platform[i][:moveTo] + "O" + platform[i][moveTo+1:]
			}
		}
	}

	// tilt south
	for i := len(lines) - 1; i >= 0; i-- {
		if i == len(lines)-1 {
			continue
		}
		for j, char := range lines[i] {
			if char == 'O' {
				moveTo := i
				for k := i + 1; k < len(lines); k++ {
					if platform[k][j] == '#' || platform[k][j] == 'O' {
						moveTo = k - 1
						break
					}
					if k == len(lines)-1 {
						moveTo = len(lines) - 1
					}
				}
				platform[i] = platform[i][:j] + "." + platform[i][j+1:]
				platform[moveTo] = platform[moveTo][:j] + "O" + platform[moveTo][j+1:]
			}
		}
	}

	// tilt east
	for i := 0; i < len(lines); i++ {
		for j := len(lines[i]) - 1; j >= 0; j-- {
			if j == len(lines[i])-1 {
				continue
			}
			if platform[i][j] == 'O' {
				moveTo := j
				for k := j + 1; k < len(lines[i]); k++ {
					if platform[i][k] == '#' || platform[i][k] == 'O' {
						moveTo = k - 1
						break
					}
					if k == len(lines[i])-1 {
						moveTo = len(lines[i]) - 1
					}
				}
				platform[i] = platform[i][:j] + "." + platform[i][j+1:]
				platform[i] = platform[i][:moveTo] + "O" + platform[i][moveTo+1:]
			}
		}
	}

	return platform
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
