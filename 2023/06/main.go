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

	fmt.Println("Multiple of ways to beat the record: ", multipleOfWaysToBeatRecord(lines))
	fmt.Println("Ways to beat the record in the longer race:", waysToBeatRecordLargeRace(lines))
}

type raceRecord struct {
	time     int
	distance int
}

func multipleOfWaysToBeatRecord(lines []string) int {
	result := 0

	raceRecords := parse(lines)
	for _, raceRecord := range raceRecords {
		ways := waysToBeatRecord(raceRecord)
		if result > 0 {
			result *= ways
		} else {
			result = ways
		}
	}

	return result
}

func waysToBeatRecord(raceRecord raceRecord) int {
	ways := 0

	for i := 0; i <= raceRecord.time; i++ {
		remaining_time := raceRecord.time - i
		distance := i * remaining_time
		if distance > raceRecord.distance {
			ways++
		}
	}

	return ways
}

func waysToBeatRecordLargeRace(lines []string) int {
	raceRecords := parse(lines)
	raceRecord := combineRaceRecords(raceRecords)
	return waysToBeatRecord(raceRecord)
}

func parse(lines []string) []raceRecord {
	raceRecords := []raceRecord{}

	for _, line := range lines {
		isTime := strings.HasPrefix(line, "Time:")
		isDistance := strings.HasPrefix(line, "Distance:")
		fields := strings.Fields(line)
		for i := 1; i < len(fields); i++ {
			val, _ := strconv.Atoi(fields[i])
			if isTime {
				raceRecords = append(raceRecords, raceRecord{val, 0})
			} else if isDistance {
				raceRecords[i-1].distance = val
			}
		}
	}

	return raceRecords
}

func combineRaceRecords(raceRecords []raceRecord) raceRecord {
	raceRecord := raceRecord{}

	for _, r := range raceRecords {
		raceRecord.time, _ = strconv.Atoi(fmt.Sprint(raceRecord.time) + fmt.Sprint(r.time))
		raceRecord.distance, _ = strconv.Atoi(fmt.Sprint(raceRecord.distance) + fmt.Sprint(r.distance))
	}

	return raceRecord
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
