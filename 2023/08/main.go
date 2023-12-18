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

	fmt.Println("Steps Part 1:", sumOfSteps(lines))
	fmt.Println("Steps Part 2:", sumOfStepsGhost(lines))

}

type node struct {
	name  string
	left  string
	right string
}
type nodesMap map[string]node

func sumOfSteps(lines []string) int {
	instructions, nodes, _ := parse(lines)
	return calcSumOfSteps(instructions, nodes, "AAA")
}

func sumOfStepsGhost(lines []string) int {
	instructions, nodes, _ := parse(lines)
	return calcSumOfStepsGhost(instructions, nodes)
}

func calcSumOfSteps(instructions string, nodes nodesMap, startNode string) int {
	sum := 0

	currentInstructions := instructions
	currentNode := startNode

	// loop until we reach ZZZ
	for currentNode != "ZZZ" {
		dir := string(currentInstructions[0])
		currentInstructions = currentInstructions[1:]

		if dir == "L" {
			currentNode = nodes[currentNode].left
		} else if dir == "R" {
			currentNode = nodes[currentNode].right
		}
		sum++

		// no more instructions? start from the beginning
		if len(currentInstructions) == 0 {
			currentInstructions = instructions
		}
	}

	return sum
}

func calcSumOfSteps2(instructions string, nodes nodesMap, startNode string) int {
	sum := 0

	current_instructions := instructions
	current_node := startNode

	// loop until we reach ZZZ
	for string(current_node[2]) != "Z" {
		dir := string(current_instructions[0])
		current_instructions = current_instructions[1:]

		if dir == "L" {
			current_node = nodes[current_node].left
		} else if dir == "R" {
			current_node = nodes[current_node].right
		}
		sum++

		// no more instructions? start from the beginning
		if len(current_instructions) == 0 {
			current_instructions = instructions
		}
	}

	return sum
}

func calcSumOfStepsGhost(instructions string, nodes nodesMap) int {
	counts := []int{}

	startingNodes := []string{}
	for _, node := range nodes {
		if strings.HasSuffix(node.name, "A") {
			startingNodes = append(startingNodes, node.name)
		}
	}

	for _, node := range startingNodes {
		counts = append(counts, calcSumOfSteps2(instructions, nodes, node))
	}

	return lcm(counts)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcmTwo(result, number)
	}
	return result
}

func lcmTwo(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func allNodesEndWithZ(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

func parse(lines []string) (string, nodesMap, string) {
	instructions := ""
	nodesMap := make(map[string]node)
	startNode := ""

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) == 1 {
			// instructions
			for _, c := range parts[0] {
				instructions += string(c)
			}
		} else if len(parts) == 4 {
			// node
			name := parts[0]
			if len(startNode) == 0 {
				startNode = name
			}
			// remove starting '(' and ',' from left
			left := parts[2][1:]
			left = left[:len(left)-1]

			// remove ')' from right
			right := parts[3][:len(parts[3])-1]

			node := node{
				name:  name,
				left:  left,
				right: right,
			}
			nodesMap[name] = node
		}
	}

	return instructions, nodesMap, startNode
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
