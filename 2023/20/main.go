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

	fmt.Println("Multiple of Low Pulses and High Pulses:", pulses(lines))
	fmt.Println("Fewest Number of Button Presses to get 1 Low Pulse to 'rx':", fewestButtonPresses(lines))
}

type module struct {
	name         string
	moduleType   string
	destinations []string
	ffState      bool            // flipFlop: on (true) or off
	cInputs      map[string]bool // conjunction inputs: high (true) or low (false)
}

type pulse struct {
	source      string
	destination string
	high        bool // true: high, false: low
}

type tracker struct {
	targets    map[string]bool
	counters   map[string]int
	iterations int
}

func pulses(lines []string) int {
	modules := parse(lines)

	sumHighPulses, sumLowPulses := 0, 0
	for i := 0; i < 1000; i++ {
		newModules, highPulses, lowPulses := pushButton(modules, nil)
		modules = newModules
		sumHighPulses += highPulses
		sumLowPulses += lowPulses
	}

	return sumHighPulses * sumLowPulses
}

func fewestButtonPresses(lines []string) int {
	modules := parse(lines)

	endModuleName := "rx"

	// This only works if the end module only has one predecessor, which is true for the input
	var predecessorModule module
	for _, module := range modules {
		for _, destination := range module.destinations {
			if destination == endModuleName {
				predecessorModule = module
				break
			}
		}
	}

	targets := make(map[string]bool)
	for cInput := range predecessorModule.cInputs {
		targets[cInput] = true
	}

	t := tracker{
		targets:    targets,
		counters:   make(map[string]int),
		iterations: 1,
	}
	for {
		newModules, _, _ := pushButton(modules, &t)
		modules = newModules
		if len(t.targets) == 0 {
			break
		}
		t.iterations++
	}

	numbers := []int{}
	for _, counter := range t.counters {
		numbers = append(numbers, counter)
	}

	return lcm(numbers)
}

func pushButton(modules []module, t *tracker) ([]module, int, int) {
	highPulses, lowPulses := 0, 0

	pulses := make([]pulse, 0)
	pulses = append(pulses, pulse{source: "button", destination: "broadcaster", high: false})

	for len(pulses) > 0 {
		p := pulses[0]
		pulses = pulses[1:]

		// count the pulse
		if p.high {
			highPulses++
		} else {
			lowPulses++
		}

		// find the module
		for i, module := range modules {
			if module.name == p.destination {
				sendHighPulse := p.high
				if module.moduleType == "%" {
					if p.high {
						// ignore high pulses
						break
					}
					modules[i].ffState = !modules[i].ffState
					sendHighPulse = modules[i].ffState
				} else if module.moduleType == "&" {
					// update memory
					module.cInputs[p.source] = p.high
					// check if all inputs are high
					allHigh := true
					for _, input := range module.cInputs {
						if !input {
							allHigh = false
							break
						}
					}
					if allHigh {
						sendHighPulse = false
					} else {
						sendHighPulse = true
					}
				}
				for _, destination := range module.destinations {
					//fmt.Println(module.name, sendHighPulse, "->", destination)
					if t != nil {
						if _, ok := t.targets[module.name]; ok && sendHighPulse {
							t.counters[module.name] = t.iterations
							delete(t.targets, module.name)
						}
						if len(t.targets) == 0 {
							break
						}
					}

					pulses = append(pulses, pulse{source: module.name, destination: destination, high: sendHighPulse})
				}
				break
			}
		}
	}

	return modules, highPulses, lowPulses
}

func parse(lines []string) []module {
	moduleConfiguration := make([]module, 0, len(lines))

	for _, line := range lines {
		m := module{}
		parts := strings.Split(line, " -> ")
		if strings.HasPrefix(parts[0], "%") {
			m.moduleType = "%"
			m.name = parts[0][1:]
		} else if strings.HasPrefix(parts[0], "&") {
			m.moduleType = "&"
			m.name = parts[0][1:]
			m.cInputs = make(map[string]bool)
		} else {
			m.name = parts[0]
		}
		m.destinations = strings.Split(parts[1], ", ")
		moduleConfiguration = append(moduleConfiguration, m)
	}

	// initialize cInputs with all incoming modules
	conjunctions := []string{}
	for _, module := range moduleConfiguration {
		if module.moduleType == "&" {
			conjunctions = append(conjunctions, module.name)
		}
	}
	for _, module := range moduleConfiguration {
		for _, destination := range module.destinations {
			for _, conjunction := range conjunctions {
				if destination == conjunction {
					for i, m := range moduleConfiguration {
						if m.name == destination {
							moduleConfiguration[i].cInputs[module.name] = false
							break
						}
					}
				}
			}
		}
	}

	return moduleConfiguration
}

func lcm(numbers []int) int {
	result := 1
	for _, number := range numbers {
		result = (result / gcd(result, number)) * number
	}
	return result
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
