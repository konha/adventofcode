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
	lines, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Lowest Location Number Part 1:", lowestLocationNum(lines))
	fmt.Println("Lowest Location Number Part 2:", lowestLocationNum2(lines))
}

func lowestLocationNum(lines []string) int {
	lowest := math.MaxInt
	seeds, maps := parse(lines)

	for _, seed := range seeds {
		soil := lookupInMap(maps["seed-to-soil"], seed)
		fertilizer := lookupInMap(maps["soil-to-fertilizer"], soil)
		water := lookupInMap(maps["fertilizer-to-water"], fertilizer)
		light := lookupInMap(maps["water-to-light"], water)
		temperature := lookupInMap(maps["light-to-temperature"], light)
		humidity := lookupInMap(maps["temperature-to-humidity"], temperature)
		location := lookupInMap(maps["humidity-to-location"], humidity)

		lowest = min(lowest, location)
	}

	return lowest
}

func lowestLocationNum2(lines []string) int {
	seedsRanges, maps := parse(lines)

	for i := 0; i < math.MaxInt; i++ {
		location := i
		humidity := reverseLookupInMap(maps["humidity-to-location"], location)
		temperature := reverseLookupInMap(maps["temperature-to-humidity"], humidity)
		light := reverseLookupInMap(maps["light-to-temperature"], temperature)
		water := reverseLookupInMap(maps["water-to-light"], light)
		fertilizer := reverseLookupInMap(maps["fertilizer-to-water"], water)
		soil := reverseLookupInMap(maps["soil-to-fertilizer"], fertilizer)
		seed := reverseLookupInMap(maps["seed-to-soil"], soil)

		for j := 0; j < len(seedsRanges); j += 2 {
			start := seedsRanges[j]
			length := seedsRanges[j+1]
			if seed >= start && seed < start+length {
				return location
			}
		}
	}

	return math.MaxInt
}

type mapLine struct {
	destRangeStart, srcRangeStart, rangeLen int
}

func lookupInMap(mapLines []mapLine, key int) int {

	for _, mapLine := range mapLines {
		if key >= mapLine.srcRangeStart && key < mapLine.srcRangeStart+mapLine.rangeLen {
			return mapLine.destRangeStart + (key - mapLine.srcRangeStart)
		}
	}

	return key
}

func reverseLookupInMap(mapLines []mapLine, key int) int {

	for _, mapLine := range mapLines {
		if key >= mapLine.destRangeStart && key < mapLine.destRangeStart+mapLine.rangeLen {
			return mapLine.srcRangeStart + (key - mapLine.destRangeStart)
		}
	}

	return key
}

func parse(lines []string) ([]int, map[string][]mapLine) {
	var seeds []int
	maps := make(map[string][]mapLine)

	currentMap := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			fields := strings.Fields(line)
			for _, field := range fields[1:] {
				var seed int
				fmt.Sscanf(field, "%d", &seed)
				seeds = append(seeds, seed)
			}

		} else if strings.HasSuffix(line, "map:") {
			currentMap = strings.TrimSuffix(line, " map:")
		} else {
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}
			var ml mapLine
			fmt.Sscanf(fields[0], "%d", &ml.destRangeStart)
			fmt.Sscanf(fields[1], "%d", &ml.srcRangeStart)
			fmt.Sscanf(fields[2], "%d", &ml.rangeLen)
			maps[currentMap] = append(maps[currentMap], ml)
		}
	}

	return seeds, maps
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
