package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2025/day5/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2025/day5/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0
	intervals := make([][2]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			var start, end int
			fmt.Sscanf(line, "%d-%d", &start, &end)
			intervals = append(intervals, [2]int{start, end})
		} else if line == "" {
			continue
		} else {
			num, _ := strconv.Atoi(line)
			for _, interval := range intervals {
				if num >= interval[0] && num <= interval[1] {
					result++
					break
				}
			}
		}
	}

	return result
}

func part2(f *os.File) int {
	intervals := make([][2]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "-") {
			break
		}
		var start, end int
		fmt.Sscanf(line, "%d-%d", &start, &end)
		intervals = append(intervals, [2]int{start, end})
	}

	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	merged := make([][2]int, 0)
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= current[1] {
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			merged = append(merged, current)
			current = intervals[i]
		}
	}
	merged = append(merged, current)

	result := 0
	for _, interval := range merged {
		result += interval[1] - interval[0] + 1
	}

	return result
}
