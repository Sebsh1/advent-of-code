package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2025/day6/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2025/day6/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0
	columns := make(map[int][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for i, part := range parts {
			columns[i] = append(columns[i], part)
		}
	}

	for i := 0; i < len(columns); i++ {
		switch columns[i][len(columns[i])-1] {
		case "*":
			prod := 1
			for _, val := range columns[i][:len(columns[i])-1] {
				num, _ := strconv.Atoi(val)
				prod *= num
			}
			result += prod
		case "+":
			sum := 0
			for _, val := range columns[i][:len(columns[i])-1] {
				num, _ := strconv.Atoi(val)
				sum += num
			}
			result += sum
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	return result
}
