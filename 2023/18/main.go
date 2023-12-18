package main

import (
	"fmt"
	"math"
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

	fmt.Println("Cubic Meters Part 1:", cubicMeters(parse1(lines)))
	fmt.Println("Cubic Meters Part 2:", cubicMeters(parse2(lines)))
}

type grid [][]string
type instruction struct {
	dir    string
	amount int
}
type point struct {
	row int
	col int
}

func cubicMeters(instructions []instruction) int {
	points := followInstructionsToPoints(instructions)
	area := shoeLace(points) + perimeter(instructions)/2 + 1
	return area
}

func calcAreaShoeLace(points []point) float64 {
	n := len(points)
	if n < 3 {
		return 0 // Not enough points to form a polygon
	}

	area := 0
	for i := 0; i < n-1; i++ {
		area += points[i].col*points[i+1].row - points[i+1].col*points[i].row
	}

	// Closing the polygon by adding the last term separately
	area += points[n-1].col*points[0].row - points[0].col*points[n-1].row

	// Taking the absolute value and dividing by 2
	return math.Abs(float64(area)) / 2.0
}

func shoeLace(points []point) int {
	var area = 0

	j := len(points) - 1

	for i := 0; i < len(points); i++ {
		area += (points[j].col + points[i].col) * (points[j].row - points[i].row)
		j = i
	}

	return int(math.Abs(float64(area) / 2.0))
}

func parse1(lines []string) []instruction {
	instructions := make([]instruction, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		color := parts[2]
		color = color[1 : len(color)-1]
		instructions = append(instructions, instruction{dir, amount})
	}
	return instructions
}

func parse2(lines []string) []instruction {
	instructions := make([]instruction, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		color := parts[2]
		color = color[1 : len(color)-1]
		// last digit of color is the direction
		dir := string(color[len(color)-1])
		// map dir to UDLR
		switch dir {
		case "0":
			dir = "R"
		case "1":
			dir = "D"
		case "2":
			dir = "L"
		case "3":
			dir = "U"
		}
		// remove last digit from color
		color = color[:len(color)-1]
		// remove the first char from color
		color = color[1:]
		// the color is a hex string, convert to int
		hex, _ := strconv.ParseInt(color, 16, 64)
		// make amount int
		amount := int(hex)

		instructions = append(instructions, instruction{dir, amount})
	}
	return instructions
}

func followInstructionsToPoints(instructions []instruction) []point {
	var points []point
	posX, posY := 0, 0
	points = append(points, point{posX, posY})
	for _, instr := range instructions {
		switch instr.dir {
		case "U":
			posY -= instr.amount
		case "D":
			posY += instr.amount
		case "L":
			posX -= instr.amount
		case "R":
			posX += instr.amount
		}
		points = append(points, point{posY, posX})
	}

	return points
}

func perimeter(instructions []instruction) int {
	var perimeter int
	for _, instr := range instructions {
		perimeter += instr.amount
	}
	return perimeter
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
