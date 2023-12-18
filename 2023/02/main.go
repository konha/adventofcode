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

	fmt.Println("Sum of valid GameIDs:", sumOfValidGameIDs(lines, 12, 13, 14))
	fmt.Println("Sum of power of minimum sets:", sumOfPowerOfMinSets(lines))
}

func sumOfValidGameIDs(lines []string, max_red, max_green, max_blue int) int {
	sum := 0

	for _, line := range lines {
		parsed, err := parseLine(line)
		if err != nil {
			return 0
		}
		if validGame(parsed, max_red, max_green, max_blue) {
			sum += parsed.id
		}
	}

	return sum
}

func sumOfPowerOfMinSets(lines []string) int {
	sum := 0

	for _, line := range lines {
		parsed, err := parseLine(line)
		if err != nil {
			return 0
		}
		minSet := minSetFromGame(parsed)
		sum += powerOfSet(minSet)
	}

	return sum
}

type set struct {
	red   int
	blue  int
	green int
}

type game struct {
	id   int
	sets []set
}

func parseLine(line string) (game, error) {
	g := game{}

	// Example: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return g, fmt.Errorf("invalid line: %s", line)
	}
	gamePart := parts[0]
	setsPart := parts[1]

	gameParts := strings.Split(gamePart, " ")
	if len(gameParts) != 2 {
		return g, fmt.Errorf("invalid game part: %s", gamePart)
	}
	gameNumber, err := strconv.Atoi(gameParts[1])
	if err != nil {
		return g, fmt.Errorf("invalid game number: %s", gameParts[1])
	}
	g.id = gameNumber

	setsParts := strings.Split(setsPart, ";")
	for _, setPart := range setsParts {
		set := set{}
		colors := strings.Split(setPart, ",")
		for _, color := range colors {
			color = strings.TrimSpace(color)
			colorParts := strings.Split(color, " ")
			if len(colorParts) != 2 {
				return g, fmt.Errorf("invalid color: %s", color)
			}
			number, err := strconv.Atoi(colorParts[0])
			if err != nil {
				return g, fmt.Errorf("invalid number: %s", colorParts[0])
			}
			color := colorParts[1]
			switch color {
			case "red":
				set.red = number
			case "blue":
				set.blue = number
			case "green":
				set.green = number
			default:
				return g, fmt.Errorf("invalid color: %s", color)
			}
		}
		g.sets = append(g.sets, set)
	}
	return g, nil
}

func validSet(set set, max_red, max_green, max_blue int) bool {
	return set.red <= max_red && set.green <= max_green && set.blue <= max_blue
}

func validGame(game game, max_red, max_green, max_blue int) bool {
	for _, set := range game.sets {
		if !validSet(set, max_red, max_green, max_blue) {
			return false
		}
	}
	return true
}

func powerOfSet(set set) int {
	return set.red * set.green * set.blue
}

func minSetFromGame(game game) set {
	min := set{}
	for _, set := range game.sets {
		if set.red > min.red {
			min.red = set.red
		}
		if set.green > min.green {
			min.green = set.green
		}
		if set.blue > min.blue {
			min.blue = set.blue
		}
	}
	return min
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
