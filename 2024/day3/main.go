package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f1, _ := os.Open("2024/day3/input.txt")
	defer f1.Close()
	fmt.Println("Part 1: ", part1(f1))

	f2, _ := os.Open("2024/day3/input.txt")
	defer f2.Close()
	fmt.Println("Part 2: ", part2(f2))
}

func part1(f *os.File) int {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	content, _ := io.ReadAll(f)

	result := 0

	matches := regex.FindAllStringSubmatch(string(content), -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		result += num1 * num2
	}

	return result
}

func part2(f *os.File) int {
	regex := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	content, _ := io.ReadAll(f)

	enabled := true
	result := 0

	matches := regex.FindAllStringSubmatch(string(content), -1)
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		}
		if match[0] == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			num1, _ := strconv.Atoi(match[2])
			num2, _ := strconv.Atoi(match[3])
			result += num1 * num2
		}
	}

	return result
}
