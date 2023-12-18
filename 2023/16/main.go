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

	fmt.Println("Energized Tiles Part 1:", countEnergizedTiles(lines))
	fmt.Println("Energizes Tiles Part 2:", maxEnergizedTiles(lines))
}

func countEnergizedTiles(lines []string) int {
	tiles := parse(lines)
	cache := map[point][]direction{}
	energized := []point{}
	energized = follow(tiles, energized, cache, point{0, 0}, right)
	return len(energized)
}

func maxEnergizedTiles(lines []string) int {
	max := 0
	tiles := parse(lines)

	// top row
	for x := range tiles[0] {
		cache := map[point][]direction{}
		energized := []point{}
		energized = follow(tiles, energized, cache, point{x, 0}, down)
		if len(energized) > max {
			max = len(energized)
		}
	}
	// bottom row
	for x := range tiles[len(tiles)-1] {
		cache := map[point][]direction{}
		energized := []point{}
		energized = follow(tiles, energized, cache, point{x, len(tiles) - 1}, up)
		if len(energized) > max {
			max = len(energized)
		}
	}
	// left column
	for y := range tiles {
		cache := map[point][]direction{}
		energized := []point{}
		energized = follow(tiles, energized, cache, point{0, y}, right)
		if len(energized) > max {
			max = len(energized)
		}
	}
	// right column
	for y := range tiles {
		cache := map[point][]direction{}
		energized := []point{}
		energized = follow(tiles, energized, cache, point{len(tiles[y]) - 1, y}, left)
		if len(energized) > max {
			max = len(energized)
		}
	}

	return max
}

func follow(tiles [][]string, energized []point, cache map[point][]direction, p point, moving direction) []point {
	if p.x < 0 || p.y < 0 || p.x >= len(tiles) || p.y >= len(tiles[p.x]) {
		return energized
	}
	tile := tiles[p.y][p.x]
	found := false
	for _, e := range energized {
		if e.x == p.x && e.y == p.y {
			found = true
			break
		}
	}
	if !found {
		energized = append(energized, p)
	}
	direction1, direction2 := tileToDirection(tile, moving)
	if direction1 != none {
		new := nextCoordinates(p, direction1)
		// see if cache has an entry for new and if it has direction1
		// if it has direction1, then we have a loop
		if _, ok := cache[new]; ok {
			for _, d := range cache[new] {
				if d == direction1 {
					return energized
				}
			}
		}
		cache[new] = append(cache[new], direction1)
		energized = follow(tiles, energized, cache, new, direction1)
	}
	if direction2 != none {
		new := nextCoordinates(p, direction2)
		// see if cache has an entry for new and if it has direction2
		// if it has direction2, then we have a loop
		if _, ok := cache[new]; ok {
			for _, d := range cache[new] {
				if d == direction2 {
					return energized
				}
			}
		}
		energized = follow(tiles, energized, cache, new, direction2)
	}
	return energized
}

func nextCoordinates(p point, moving direction) point {
	switch moving {
	case right:
		return point{p.x + 1, p.y}
	case down:
		return point{p.x, p.y + 1}
	case left:
		return point{p.x - 1, p.y}
	case up:
		return point{p.x, p.y - 1}
	}
	return point{p.x, p.y}
}

type point struct {
	x, y int
}

type direction int

const (
	none direction = iota
	right
	down
	left
	up
)

func tileToDirection(tile string, moving direction) (direction, direction) {
	switch tile {
	case ".":
		return moving, none
	case "\\":
		switch moving {
		case right:
			return down, none
		case down:
			return right, none
		case left:
			return up, none
		case up:
			return left, none
		}
	case "/":
		switch moving {
		case right:
			return up, none
		case down:
			return left, none
		case left:
			return down, none
		case up:
			return right, none
		}
	case "|":
		switch moving {
		case right:
			return up, down
		case down:
			return down, none
		case left:
			return up, down
		case up:
			return up, none
		}
	case "-":
		switch moving {
		case right:
			return right, none
		case down:
			return left, right
		case left:
			return left, none
		case up:
			return left, right
		}
	}
	return none, none
}

func parse(lines []string) [][]string {
	var result [][]string
	for _, line := range lines {
		result = append(result, strings.Split(line, ""))
	}
	return result
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
