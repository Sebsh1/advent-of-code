package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	None Direction = iota
	Increasing
	Decreasing
)

func main() {
	f1, _ := os.Open("2024/day2/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2024/day2/input.txt")
	defer f2.Close()
	fmt.Println("Part 2: ", part2(f2))
}

func part1(f *os.File) int {
	safeReports := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levelsStr := strings.Fields(scanner.Text())
		levels := make([]int, 0, len(levelsStr))

		for _, levelStr := range levelsStr {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		if isSafe(levels) {
			safeReports++
		}
	}

	return safeReports
}

func part2(f *os.File) int {
	safeReports := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levelsStr := strings.Fields(scanner.Text())
		levels := make([]int, 0, len(levelsStr))

		for _, levelStr := range levelsStr {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		if isSafe(levels) {
			safeReports++
			continue
		}

		for skipIdx := 0; skipIdx < len(levels); skipIdx++ {
			permutation := make([]int, 0, len(levels)-1)
			for i, level := range levels {
				if i != skipIdx {
					permutation = append(permutation, level)
				}
			}

			if isSafe(permutation) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func isSafe(levels []int) bool {
	direction := None

	for i, level := range levels {
		if i == 0 {
			continue
		}

		if i == 1 {
			if levels[i-1] < level {
				direction = Increasing
			} else {
				direction = Decreasing
			}
		}

		diff := int(math.Abs(float64(level - levels[i-1])))
		if diff < 1 || diff > 3 {
			return false
		}

		if direction == Increasing && levels[i-1] > level {
			return false
		}

		if direction == Decreasing && levels[i-1] < level {
			return false
		}
	}

	return true
}
