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

	fmt.Println("Sum of Points:", sumOfPoints(lines))
	fmt.Println("Sum of ScratchCards", sumOfScratchCards(lines))
}

func sumOfPoints(line []string) int {
	sum := 0

	for _, l := range line {
		c := parseLine(l)
		sum += points(winningNumbersOnCard(c))
	}

	return sum
}

type card struct {
	id      int
	winning []int
	have    []int
}

func parseLine(line string) card {
	c := card{}

	// Example: Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

	partID := strings.Split(line, ":")[0]
	partIDa := strings.Split(partID, " ")
	partID = partIDa[len(partIDa)-1]
	c.id, _ = strconv.Atoi(partID)

	partWinning := strings.Split(line, "|")[0]
	partWinning = strings.Split(partWinning, ":")[1]
	partWinningNums := strings.Fields(partWinning)
	winning := make([]int, len(partWinningNums))
	for i, p := range partWinningNums {
		num, _ := strconv.Atoi(p)
		winning[i] = num
	}
	c.winning = winning

	partHave := strings.Split(line, "|")[1]
	partHave = strings.TrimSpace(partHave)
	partHaveNums := strings.Fields(partHave)
	have := make([]int, len(partHaveNums))
	for i, p := range partHaveNums {
		num, _ := strconv.Atoi(p)
		have[i] = num

	}
	c.have = have

	return c
}

func winningNumbersOnCard(card card) int {
	winners := 0

	for _, h := range card.have {
		for _, w := range card.winning {
			if h == w {
				winners++
			}
		}
	}

	return winners
}

func points(winningNumbersOnCard int) int {
	if winningNumbersOnCard == 0 {
		return 0
	}

	points := 1
	for i := 1; i < winningNumbersOnCard; i++ {
		points *= 2
	}

	return points
}

func sumOfScratchCards(lines []string) int {

	wonCopies := make(map[int]int)
	for _, l := range lines {
		c := parseLine(l)
		winningNumbersOnCard := winningNumbersOnCard(c)
		if winningNumbersOnCard > 0 {
			copies := wonCopies[c.id] + 1
			for i := 1; i <= winningNumbersOnCard; i++ {
				wonCopies[c.id+i] += copies
			}
		}
	}

	sum := 0

	for _, l := range lines {
		sum += 1
		c := parseLine(l)
		if wonCopies[c.id] > 0 {
			sum += wonCopies[c.id]
		}
	}

	return sum
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
