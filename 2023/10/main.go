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

	fmt.Println("Steps to farthest Point:", stepsToFarthestPoint(lines))
	fmt.Println("Enclosed Tiles: ", enclosedTiles(lines))
}

type tileMap [][]string
type direction int

const (
	north direction = iota + 1
	south
	west
	east
)

func stepsToFarthestPoint(lines []string) int {
	steps := 0

	pipes := parse(lines)
	xStart, yStart := findStart(pipes)
	directionStart := possibleMove(pipes, xStart, yStart)
	if directionStart == 0 {
		return 0
	}

	x, y := xStart, yStart
	dir := directionStart
	for {
		x, y = followDirection(pipes, x, y, dir)
		steps++

		// are we at the start again?
		if x == xStart && y == yStart {
			return steps / 2
		}

		// where can we go from here?
		dir1, dir2 := symbolToDirections(pipes[y][x])
		if dir1 == opositeDirection(dir) {
			dir = dir2
		} else {
			dir = dir1
		}
	}
}

func symbolToDirections(symbol string) (direction, direction) {
	switch symbol {
	case "|":
		return north, south
	case "-":
		return west, east
	case "L":
		return north, east
	case "J":
		return north, west
	case "7":
		return south, west
	case "F":
		return south, east
	default:
		return 0, 0
	}
}

func opositeDirection(dir direction) direction {
	switch dir {
	case north:
		return south
	case south:
		return north
	case west:
		return east
	case east:
		return west
	default:
		return 0
	}
}

func followDirection(pipes tileMap, x, y int, dir direction) (int, int) {
	switch dir {
	case north:
		return x, y - 1
	case south:
		return x, y + 1
	case west:
		return x - 1, y
	case east:
		return x + 1, y
	default:
		return x, y
	}
}

func findStart(pipes tileMap) (int, int) {
	for y, row := range pipes {
		for x, symbol := range row {
			if symbol == "S" {
				return x, y
			}
		}
	}
	return 0, 0
}

func possibleMove(pipes tileMap, x, y int) direction {

	moves := [][]int{
		{x, y - 1}, // north
		{x + 1, y}, // east
		{x, y + 1}, // south
		{x - 1, y}, // west
	}

	for _, move := range moves {
		if move[0] < 0 || move[1] < 0 {
			continue
		}
		if move[0] >= len(pipes[0]) || move[1] >= len(pipes) {
			continue
		}
		symbol := pipes[move[1]][move[0]]
		direction1, direction2 := symbolToDirections(symbol)
		if direction1 == 0 || direction2 == 0 {
			continue
		}
		if move[0] < x && (direction1 == east || direction2 == east) {
			return west
		} else if move[0] > x && (direction1 == west || direction2 == west) {
			return east
		} else if move[1] < y && (direction1 == south || direction2 == south) {
			return north
		} else if move[1] > y && (direction1 == north || direction2 == north) {
			return south
		}
	}

	return 0
}

func enclosedTiles(lines []string) int {
	tileMap := parse(lines)
	pipeTiles := pipeTiles(tileMap)

	// scan for enclosed tiles
	count := 0

	for _, row := range pipeTiles {
		prev_corner := ""
		crossings := 0
		for _, symbol := range row {
			if symbol != "" {
				switch symbol {
				case "|":
					crossings += 1
				case "7":
					if prev_corner == "L" || prev_corner == "S" {
						crossings += 1
					}
				case "J":
					if prev_corner == "F" || prev_corner == "S" {
						crossings += 1
					}
				case "S":
					crossings += 1
				}
				if symbol != "-" {
					prev_corner = symbol
				}
				//fmt.Print("_")
			} else {
				if crossings%2 == 1 {
					//fmt.Print("X")
					count += 1
				} else {
					//fmt.Print("_")
				}
			}
		}
		//fmt.Println()
	}

	return count

}

func pipeTiles(tm tileMap) tileMap {
	pipeTiles := make(tileMap, len(tm))
	for i := range pipeTiles {
		pipeTiles[i] = make([]string, len(tm[i]))
	}

	xStart, yStart := findStart(tm)
	directionStart := possibleMove(tm, xStart, yStart)
	if directionStart == 0 {
		return nil
	}

	x, y := xStart, yStart
	dir := directionStart
	for {
		pipeTiles[y][x] = tm[y][x]
		x, y = followDirection(tm, x, y, dir)

		// are we at the start again?
		if x == xStart && y == yStart {
			break
		}

		// where can we go from here?
		dir1, dir2 := symbolToDirections(tm[y][x])
		if dir1 == opositeDirection(dir) {
			dir = dir2
		} else {
			dir = dir1
		}
	}

	return pipeTiles
}

func parse(lines []string) tileMap {
	pipes := make(tileMap, len(lines))
	for i, line := range lines {
		for _, char := range line {
			pipes[i] = append(pipes[i], string(char))
		}
	}
	return pipes
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
