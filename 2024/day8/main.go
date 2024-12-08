package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day8/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	start = time.Now()
	result = part2()
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

var grid = [][]rune{}
var antennas = []node{}

type node struct {
	row, i int
	freq   rune
}

func nodeString(n node) string {
	return fmt.Sprintf("%d,%d", n.row, n.i)
}

func getNextNode(a1, a2 node) (bool, *node) {
	if a1 == a2 || a1.freq != a2.freq {
		return false, nil
	}

	nextNode := &node{a2.row - (a1.row - a2.row), a2.i - (a1.i - a2.i), a1.freq}

	if nextNode.row < 0 || nextNode.row >= len(grid) || nextNode.i < 0 || nextNode.i >= len(grid[nextNode.row]) {
		return false, nil
	}

	return true, nextNode
}

func part1(f *os.File) int {
	result := 0

	antiNodeLocations := make(map[string]struct{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)
		for _, r := range line {
			row = append(row, r)
			if r != '.' {
				antennas = append(antennas, node{len(grid), len(row) - 1, r})
			}
		}
		grid = append(grid, row)
	}

	for _, a1 := range antennas {
		for _, a2 := range antennas {
			valid, antiNode := getNextNode(a1, a2)
			if !valid {
				continue
			}
			antiNodeLocations[nodeString(*antiNode)] = struct{}{}
		}
	}

	for range antiNodeLocations {
		result++
	}

	return result
}

func part2() int {
	result := 0
	antiNodeLocations := make(map[string]struct{})

	for _, a1 := range antennas {
		for _, a2 := range antennas {
			currNode := a2
			nextNode := a1
			for {
				if currNode != nextNode || currNode.freq == nextNode.freq {
					antiNodeLocations[nodeString(currNode)] = struct{}{}
					antiNodeLocations[nodeString(nextNode)] = struct{}{}
				}

				valid, antiNode := getNextNode(currNode, nextNode)
				if !valid {
					break
				}
				antiNodeLocations[nodeString(*antiNode)] = struct{}{}

				currNode, nextNode = nextNode, *antiNode
			}
		}
	}

	for range antiNodeLocations {
		result++
	}

	return result
}
