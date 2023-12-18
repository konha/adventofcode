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
	filename := os.Args[1]

	lines, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sum of Results:", sumOfResults(lines))
	fmt.Println("Focussing Power: ", focusingPower(lines))
}

func sumOfResults(lines []string) int {
	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, ",")
		for _, part := range parts {
			hash := calcHash(part)
			sum += hash
		}
	}

	return sum
}

func focusingPower(lines []string) int {
	boxes := make(map[int][]string, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = make([]string, 0)
	}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		for _, part := range parts {

			if strings.Contains(part, "-") {
				label := strings.Split(part, "-")[0]
				box := calcHash(label)
				// remove lens with label from box
				for i, content := range boxes[box] {
					if strings.Split(content, " ")[0] == label {
						boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
						break
					}
				}
			} else if strings.Contains(part, "=") {
				label := strings.Split(part, "=")[0]
				box := calcHash(label)
				focalLength := strings.Split(part, "=")[1]
				// if there is a lens with label in box, replace it
				found := false
				for i, content := range boxes[box] {
					if strings.Split(content, " ")[0] == label {
						boxes[box][i] = label + " " + focalLength
						found = true
						break
					}
				}
				// if not, add it to the end
				if !found {
					boxes[box] = append(boxes[box], label+" "+focalLength)
				}
			}
		}
	}

	sum := 0
	for i := 0; i < 256; i++ {
		box := boxes[i]
		for j, lens := range box {
			focalLength := strings.Split(lens, " ")[1]
			fl, _ := strconv.Atoi(focalLength)
			val := (i + 1) * (j + 1) * fl
			sum += val
		}
	}
	return sum
}

func calcHash(s string) int {
	currentValue := 0

	for _, c := range s {
		ascii := int(c)
		currentValue += ascii
		currentValue *= 17
		currentValue = currentValue % 256
	}

	return currentValue
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
