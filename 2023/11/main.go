package main

import (
	"fmt"
	"math"
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

	fmt.Println("Sum of Shortest Paths Part 1:", sumOfShortestPaths(lines))
	fmt.Println("Sum of Shortest Paths Part 2:", sumOfShortestPathsHuge(lines))

}

func sumOfShortestPaths(lines []string) int {
	galaxies := parse(lines)
	galaxies = expand(galaxies, 2)
	pairs := galaxyPairs(galaxies)

	sum := 0
	for _, pair := range pairs {
		sum += shortestPath(galaxies, pair)
	}
	return sum
}

func sumOfShortestPathsHuge(lines []string) int {
	galaxies := parse(lines)

	expandedRows := make([]int, 0)
	expandedCols := make([]int, 0)

	for i, row := range galaxies {
		if rowContainsNoGalaxies(row) {
			expandedRows = append(expandedRows, i)
		}
	}
	for i := range galaxies[0] {
		if colContainsNoGalaxies(galaxies, i) {
			expandedCols = append(expandedCols, i)
		}
	}

	return calcSum(galaxies, expandedRows, expandedCols, 1000000)
}

func calcSum(galaxies [][]int, expandedRows, expandedCols []int, expansion int) int {
	g := []Point{}

	for i, row := range galaxies {
		for j := range row {
			if galaxies[i][j] != 0 {
				g = append(g, Point{i, j})
			}
		}
	}

	var count float64

	for _, p := range g {
		for _, q := range g {
			if p == q {
				continue
			}
			x_diff := expandedColsBetween(expandedCols, p.y, q.y, expansion)
			y_diff := expandedRowsBetween(expandedRows, p.x, q.x, expansion)

			if p.x > q.x {
				y_diff = -y_diff
			}
			if p.y > q.y {
				x_diff = -x_diff
			}

			x := math.Abs(float64((q.x + y_diff) - p.x))
			y := math.Abs(float64((q.y + x_diff) - p.y))
			count += x + y
		}
	}

	return int(count / 2)

}

func expandedRowsBetween(expandedRows []int, y1, y2 int, expansion int) int {
	yStart, yEnd := y1, y2
	if y2 < y1 {
		yStart, yEnd = y2, y1
	}

	count := 0
	for _, row := range expandedRows {
		if row > yStart && row < yEnd {
			count++
		}
	}
	return (count * expansion) - count
}

func expandedColsBetween(expandedCols []int, x1, x2 int, expansion int) int {
	xStart, xEnd := x1, x2
	if x2 < x1 {
		xStart, xEnd = x2, x1
	}

	count := 0
	for _, col := range expandedCols {
		if col > xStart && col < xEnd {
			count++
		}
	}
	return (count * expansion) - count
}

type Point struct {
	x, y int
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPath(galaxies [][]int, pair [2]int) int {
	startX, startY := findGalaxy(galaxies, pair[0])
	endX, endY := findGalaxy(galaxies, pair[1])

	visited := make([][]bool, len(galaxies))
	for i := range visited {
		visited[i] = make([]bool, len(galaxies[0]))
	}

	queue := []Point{{startX, startY}}
	visited[startX][startY] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]

			if point.x == endX && point.y == endY {
				return steps
			}

			for _, direction := range directions {
				newX, newY := point.x+direction.x, point.y+direction.y
				if newX >= 0 && newX < len(galaxies) && newY >= 0 && newY < len(galaxies[0]) && !visited[newX][newY] {
					queue = append(queue, Point{newX, newY})
					visited[newX][newY] = true
				}
			}
		}
		steps++
	}

	return -1
}

func findGalaxy(galaxies [][]int, galaxy int) (int, int) {
	for i, row := range galaxies {
		for j, num := range row {
			if num == galaxy {
				return i, j
			}
		}
	}
	return -1, -1
}

func galaxyPairs(galaxies [][]int) [][2]int {
	flattened := make([]int, 0)
	for _, row := range galaxies {
		for _, num := range row {
			if num != 0 {
				flattened = append(flattened, num)
			}
		}
	}

	pairs := make([][2]int, 0)
	for i := 0; i < len(flattened); i++ {
		for j := i + 1; j < len(flattened); j++ {
			pairs = append(pairs, [2]int{flattened[i], flattened[j]})
		}
	}

	return pairs
}

func expand(galaxies [][]int, by int) [][]int {
	expandedByRow := make([][]int, 0)

	for _, row := range galaxies {
		expandedByRow = append(expandedByRow, row)
		if rowContainsNoGalaxies(row) {
			for i := 0; i < by-1; i++ {
				expandedByRow = append(expandedByRow, make([]int, len(row)))
			}
		}
	}

	expandedByCol := make([][]int, 0)

	for i := 0; i < len(expandedByRow); i++ {
		expandedByCol = append(expandedByCol, make([]int, 0))
	}

	for i := range expandedByRow[0] {
		for j, row := range expandedByRow {
			expandedByCol[j] = append(expandedByCol[j], row[i])
		}
		if colContainsNoGalaxies(expandedByRow, i) {
			for k := range expandedByRow {
				expandedByCol[k] = append(expandedByCol[k], 0)
			}
		}
	}

	return expandedByCol
}

func rowContainsNoGalaxies(row []int) bool {
	for _, i := range row {
		if i != 0 {
			return false
		}
	}
	return true
}

func colContainsNoGalaxies(galaxies [][]int, col int) bool {
	for _, row := range galaxies {
		if row[col] != 0 {
			return false
		}
	}
	return true
}

func parse(lines []string) [][]int {
	galaxies := make([][]int, len(lines))

	count := 1
	for i, line := range lines {
		for _, c := range line {
			switch c {
			case '.':
				galaxies[i] = append(galaxies[i], 0)
			case '#':
				galaxies[i] = append(galaxies[i], count)
				count++
			}
		}
	}

	return galaxies
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
