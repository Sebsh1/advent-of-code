package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2024/day5/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2024/day5/input.txt")
	defer f1.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	rulesPart := true
	rules := make(map[string][]string)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		good := true
		line := scanner.Text()
		if rulesPart {
			if line == "" {
				rulesPart = false
				continue
			}
			rule := strings.Split(line, "|")
			rules[rule[0]] = append(rules[rule[0]], rule[1])
			continue
		}

		order := strings.Split(line, ",")
		orderSoFar := make([]string, 0)
		for _, page := range order {
			for _, laterPage := range rules[page] {
				if slices.Contains(orderSoFar, laterPage) {
					good = false
				}
			}
			orderSoFar = append(orderSoFar, page)
		}
		if good {
			middelPageInt, _ := strconv.Atoi(order[len(order)/2])
			result += middelPageInt
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	rulesPart := true
	rules := make(map[string][]string)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		startedBad := false
		stillBad := false
		line := scanner.Text()
		if rulesPart {
			if line == "" {
				rulesPart = false
				continue
			}
			rule := strings.Split(line, "|")
			rules[rule[0]] = append(rules[rule[0]], rule[1])
			continue
		}

		order := strings.Split(line, ",")
	retry:
		orderSoFar := make([]string, 0, len(order))
		for i, _ := range order {
			for _, laterPage := range rules[order[i]] {
				if slices.Contains(orderSoFar, laterPage) {
					startedBad = true
					stillBad = true
					j := indexOf(order, laterPage)
					order[i], order[j] = order[j], order[i]
					if stillBad {
						goto retry
					}
				}
			}
			orderSoFar = append(orderSoFar, order[i])
		}
		if startedBad {
			middelPageInt, _ := strconv.Atoi(order[len(order)/2])
			result += middelPageInt
		}
	}

	return result
}

func indexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
