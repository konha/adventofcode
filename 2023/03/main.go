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
	lines, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sum of Part Numbers:", sumOfPartNumbers(lines))
	fmt.Println("Sum of Gear Ratios:", sumOfGearRatios(lines))
}

func isSymbol(char rune) bool {
	return char != '.' && !isDigit(char)
}

func isGearSymbol(char rune) bool {
	return char == '*'
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func sumOfPartNumbers(lines []string) int {
	sum := 0

	for i, line := range lines {
		current_number := ""
		for j, char := range line {
			if isDigit(char) {
				current_number += string(char)
				if j == len(line)-1 {
					length := len(current_number)
					n, err := strconv.Atoi(current_number)
					if err != nil {
						fmt.Println(err)
						return 0
					}
					if hasAdjacentSymbol(lines, i, j-length+1, length) {
						sum += n
					}
					current_number = ""
				}
			} else {
				if current_number != "" {
					length := len(current_number)
					n, err := strconv.Atoi(current_number)
					if err != nil {
						fmt.Println(err)
						return 0
					}
					if hasAdjacentSymbol(lines, i, j-length, length) {
						sum += n
					}
					current_number = ""
				}
			}
		}
	}

	return sum
}

func hasAdjacentSymbol(lines []string, line, pos, length int) bool {
	// check left
	if pos-1 >= 0 && isSymbol(rune(lines[line][pos-1])) {
		return true
	}

	// check right
	if pos+length < len(lines[line]) && isSymbol(rune(lines[line][pos+length])) {
		return true
	}

	// check up and down
	for lineNum := line - 1; lineNum <= line+1; lineNum += 2 {
		if lineNum < 0 || lineNum >= len(lines) {
			continue
		}
		for i := pos - 1; i <= pos+length; i++ {
			if (i >= 0) && (i < len(lines[lineNum])) {
				if isSymbol(rune(lines[lineNum][i])) {
					return true
				}
			}
		}
	}

	return false
}

func sumOfGearRatios(lines []string) int {
	sum := 0

	for i, line := range lines {
		for j, char := range line {
			if isGearSymbol(char) {
				nums := adjacentNumbers(lines, i, j)
				if len(nums) == 2 {
					gearRatio := nums[0] * nums[1]
					sum += gearRatio
				}
			}
		}
	}

	return sum
}

func adjacentNumbers(lines []string, line, pos int) []int {
	nums := []int{}

	// check left
	if pos-1 >= 0 && isDigit(rune(lines[line][pos-1])) {
		num := ""
		for i := pos - 1; i >= 0; i-- {
			if isDigit(rune(lines[line][i])) {
				// add in front of num
				num = string(lines[line][i]) + num
			} else {
				break
			}
		}
		if num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			nums = append(nums, n)
		}
	}

	// check right
	if pos+1 < len(lines[line]) && isDigit(rune(lines[line][pos+1])) {
		num := ""
		for i := pos + 1; i < len(lines[line]); i++ {
			if isDigit(rune(lines[line][i])) {
				// add to num
				num += string(lines[line][i])
			} else {
				break
			}
		}
		if num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			nums = append(nums, n)
		}
	}

	// check up and down
	for lineNum := line - 1; lineNum <= line+1; lineNum += 2 {
		if lineNum < 0 || lineNum >= len(lines) {
			continue
		}
		for i := pos - 1; i <= pos+1; i++ {
			if (i >= 0) && (i < len(lines[lineNum])) && isDigit(rune(lines[lineNum][i])) {
				// go left until we hit a symbol or the start of the line
				start := i
				for j := i; j >= 0; j-- {
					if isDigit(rune(lines[lineNum][j])) {
						start = j
						if j == 0 {
							break
						}
					} else {
						break
					}
				}
				// go right until we hit a symbol or the end of the line
				num := ""
				for j := start; j < len(lines[lineNum]); j++ {
					if isDigit(rune(lines[lineNum][j])) {
						num += string(lines[lineNum][j])
					} else {
						break
					}
				}
				if num != "" {
					n, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println(err)
						return nil
					}
					nums = append(nums, n)

					if start+len(num) > pos {
						break
					}
				}
			}
		}

	}

	return nums
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
