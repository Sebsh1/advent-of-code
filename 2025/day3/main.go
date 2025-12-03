package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f1, _ := os.Open("2025/day3/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2025/day3/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		max := 0
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				if i == j {
					continue
				}
				num, _ := strconv.Atoi(string(line[i]) + string(line[j]))
				if num > max {
					max = num
				}
			}
		}

		result += max
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digits := ""
		startIdx := 0

		for len(digits) < 12 {
			max := byte('0')
			idx := startIdx
			remaining := 12 - len(digits)
			endIdx := len(line) - remaining
			for i := startIdx; i <= endIdx; i++ {
				if line[i] > max {
					max = line[i]
					idx = i
				}
			}
			digits += string(max)
			startIdx = idx + 1
		}

		num, _ := strconv.Atoi(digits)
		result += num
	}

	return result
}
