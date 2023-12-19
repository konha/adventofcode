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

	fmt.Println("Sum of Rating Numbers of Accepted Parts:", sumOfRatingNumbersOfAcceptedParts(lines))
	fmt.Println("Distinct Combinations of Ratings:", distinctCombinationsOfRatings(lines))
}

func sumOfRatingNumbersOfAcceptedParts(lines []string) int {
	workflows, parts := parse(lines)

	sum := 0
	for _, part := range parts {
		if acceptedAfterWorkflows(workflows, part) {
			sum += sumOfRatingNumbers(part)
		}
	}
	return sum
}

func distinctCombinationsOfRatings(lines []string) int {
	workflows, _ := parse(lines)

	ranges := partRatingRanges{
		x: partRatingRange{min: 1, max: 4000},
		m: partRatingRange{min: 1, max: 4000},
		a: partRatingRange{min: 1, max: 4000},
		s: partRatingRange{min: 1, max: 4000},
	}
	rangesAccepted := rangesAcceptedAfterWorkflows(workflows, ranges)

	combinations := 0
	for _, ra := range rangesAccepted {
		combinations += (ra.x.max - ra.x.min + 1) * (ra.m.max - ra.m.min + 1) * (ra.a.max - ra.a.min + 1) * (ra.s.max - ra.s.min + 1)
	}
	return combinations
}

func sumOfRatingNumbers(part part) int {
	sum := 0
	for _, r := range part.rating {
		sum += r
	}
	return sum
}

func acceptedAfterWorkflows(workflows []workflow, part part) bool {
	currentWorkflow := "in"
	for {
		result := processWorkflow(workflows, part, currentWorkflow)
		if result == "R" {
			return false
		}
		if result == "A" {
			return true
		}
		currentWorkflow = result
	}
}

func processWorkflow(workflows []workflow, part part, wfName string) string {
	workflow := getWorkflow(workflows, wfName)
	for _, rule := range workflow.rules {
		if rule.cond.cat == 0 {
			return rule.workflow
		}
		if rule.cond.comp == lt {
			if part.rating[rule.cond.cat] < rule.cond.val {
				return rule.workflow
			}
		} else if rule.cond.comp == mt {
			if part.rating[rule.cond.cat] > rule.cond.val {
				return rule.workflow
			}
		}
	}

	return workflow.name
}

func getWorkflow(workflows []workflow, wfName string) workflow {
	for _, wf := range workflows {
		if wf.name == wfName {
			return wf
		}
	}
	return workflow{rules: []rule{{workflow: "R"}}}
}

func rangesAcceptedAfterWorkflows(workflows []workflow, ranges partRatingRanges) []partRatingRanges {
	accepted := make([]partRatingRanges, 0)

	checking := make([]partRatingRanges, 0)
	checking = append(checking, ranges)

	for len(checking) > 0 {
		currentWorkflow := "in"
		done := false
		current := checking[0]
		for !done {
			if currentWorkflow == "R" {
				done = true
			}
			if currentWorkflow == "A" {
				accepted = append(accepted, current)
				done = true
			}
			nextWorkflow, nextRanges := processWorkflowForRanges(workflows, current, currentWorkflow)
			if len(nextRanges) > 1 {
				checking = append(checking, nextRanges[1])
			}
			current = nextRanges[0]
			currentWorkflow = nextWorkflow
		}
		checking = checking[1:]
	}

	return accepted
}

func processWorkflowForRanges(workflows []workflow, partRanges partRatingRanges, workflowName string) (string, []partRatingRanges) {
	nextWorkflow := workflowName
	nextRanges := partRanges

	for _, rule := range getWorkflow(workflows, workflowName).rules {
		if rule.cond.cat == 0 {
			nextWorkflow = rule.workflow
		} else {
			switch rule.cond.cat {
			case x:
				if rule.cond.comp == lt {
					if partRanges.x.max < rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.x.max > rule.cond.val && partRanges.x.min < rule.cond.val {
						nextRanges.x.min = rule.cond.val
						nextRanges.x.max = partRanges.x.max
						partRanges.x.max = rule.cond.val - 1
						nextWorkflow = rule.workflow
					}
				} else if rule.cond.comp == mt {
					if partRanges.x.min > rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.x.min < rule.cond.val && partRanges.x.max > rule.cond.val {
						nextRanges.x.min = partRanges.x.min
						nextRanges.x.max = rule.cond.val
						partRanges.x.min = rule.cond.val + 1
						nextWorkflow = rule.workflow
					}
				}
			case m:
				if rule.cond.comp == lt {
					if partRanges.m.max < rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.m.max > rule.cond.val && partRanges.m.min < rule.cond.val {
						nextRanges.m.min = rule.cond.val
						nextRanges.m.max = partRanges.m.max
						partRanges.m.max = rule.cond.val - 1
						nextWorkflow = rule.workflow
					}
				} else if rule.cond.comp == mt {
					if partRanges.m.min > rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.m.min < rule.cond.val && partRanges.m.max > rule.cond.val {
						nextRanges.m.min = partRanges.m.min
						nextRanges.m.max = rule.cond.val
						partRanges.m.min = rule.cond.val + 1
						nextWorkflow = rule.workflow
					}
				}
			case a:
				if rule.cond.comp == lt {
					if partRanges.a.max < rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.a.max > rule.cond.val && partRanges.a.min < rule.cond.val {
						nextRanges.a.min = rule.cond.val
						nextRanges.a.max = partRanges.a.max
						partRanges.a.max = rule.cond.val - 1
						nextWorkflow = rule.workflow
					}
				} else if rule.cond.comp == mt {
					if partRanges.a.min > rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.a.min < rule.cond.val && partRanges.a.max > rule.cond.val {
						nextRanges.a.min = partRanges.a.min
						nextRanges.a.max = rule.cond.val
						partRanges.a.min = rule.cond.val + 1
						nextWorkflow = rule.workflow
					}
				}
			case s:
				if rule.cond.comp == lt {
					if partRanges.s.max < rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.s.max > rule.cond.val && partRanges.s.min < rule.cond.val {
						nextRanges.s.min = rule.cond.val
						nextRanges.s.max = partRanges.s.max
						partRanges.s.max = rule.cond.val - 1
						nextWorkflow = rule.workflow
					}
				} else if rule.cond.comp == mt {
					if partRanges.s.min > rule.cond.val {
						nextWorkflow = rule.workflow
					} else if partRanges.s.min < rule.cond.val && partRanges.s.max > rule.cond.val {
						nextRanges.s.min = partRanges.s.min
						nextRanges.s.max = rule.cond.val
						partRanges.s.min = rule.cond.val + 1
						nextWorkflow = rule.workflow
					}
				}
			}
		}
		if nextWorkflow != workflowName {
			break
		}
	}

	res := []partRatingRanges{partRanges}

	if partRanges != nextRanges {
		res = append(res, nextRanges)
	}

	return nextWorkflow, res
}

type workflow struct {
	name  string
	rules []rule
}
type rule struct {
	cond     condition
	workflow string
}
type condition struct {
	cat  category
	comp comparison
	val  int
}
type category int

const (
	x category = iota + 1
	m
	a
	s
)

type comparison int

const (
	lt comparison = iota
	mt
)

type part struct {
	rating map[category]int
}

type partRatingRange struct {
	min, max int
}
type partRatingRanges struct {
	x, m, a, s partRatingRange
}

func parse(lines []string) ([]workflow, []part) {
	workflows := make([]workflow, 0)
	parts := make([]part, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "{") {
			parts = append(parts, parsePart(line))
		} else {
			workflows = append(workflows, parseWorkflow(line))
		}
	}
	return workflows, parts
}

func parseWorkflow(line string) workflow {
	w := workflow{}
	w.rules = make([]rule, 0)

	w.name = strings.Split(line, "{")[0]
	rulesPart := strings.Split(strings.Split(line, "{")[1], "}")[0]
	for _, r := range strings.Split(rulesPart, ",") {
		if strings.Contains(r, ":") {
			wfNamePart := strings.Split(r, ":")[1]
			wfCondPart := strings.Split(r, ":")[0]
			if strings.Contains(wfCondPart, "<") {
				catPart := strings.Split(wfCondPart, "<")[0]
				comp := lt
				val, _ := strconv.Atoi(strings.Split(wfCondPart, "<")[1])
				w.rules = append(w.rules, rule{workflow: wfNamePart, cond: condition{cat: stringToCategory(catPart), comp: comp, val: val}})
			} else if strings.Contains(wfCondPart, ">") {
				catPart := strings.Split(wfCondPart, ">")[0]
				comp := mt
				val, _ := strconv.Atoi(strings.Split(wfCondPart, ">")[1])
				w.rules = append(w.rules, rule{workflow: wfNamePart, cond: condition{cat: stringToCategory(catPart), comp: comp, val: val}})
			}
		} else {
			w.rules = append(w.rules, rule{workflow: r})
		}
	}

	return w
}

func stringToCategory(char string) category {
	return map[string]category{"x": x, "m": m, "a": a, "s": s}[char]
}

func parsePart(line string) part {
	part := part{}
	part.rating = make(map[category]int)

	line = strings.TrimPrefix(line, "{")
	line = strings.TrimSuffix(line, "}")

	ps := strings.Split(line, ",")
	for _, p := range ps {
		cat := strings.Split(p, "=")[0]
		val, _ := strconv.Atoi(strings.Split(p, "=")[1])
		part.rating[stringToCategory(cat)] = val
	}

	return part
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
