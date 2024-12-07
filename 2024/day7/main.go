package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day7/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	f2, _ := os.Open("2024/day7/input.txt")
	defer f2.Close()
	start = time.Now()
	result = part2(f2)
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

func part1(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		numbers := make([]int, 0)
		for _, numberStr := range strings.Split(parts[1], " ") {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}
		if rec1(0, target, numbers) {
			result += target
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		numbers := make([]int, 0)
		for _, numberStr := range strings.Split(parts[1], " ") {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}
		if rec2(0, target, numbers) {
			result += target
		}
	}

	return result
}

func rec1(start, target int, numbers []int) bool {
	if len(numbers) == 0 && start == target {
		return true
	}

	if len(numbers) == 0 {
		return false
	}

	if rec1(start+numbers[0], target, numbers[1:]) || rec1(start*numbers[0], target, numbers[1:]) {
		return true
	}

	return false
}

func rec2(start, target int, numbers []int) bool {
	if len(numbers) == 0 && start == target {
		return true
	}

	if len(numbers) == 0 {
		return false
	}

	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", start, numbers[0]))

	if rec2(start+numbers[0], target, numbers[1:]) ||
		rec2(start*numbers[0], target, numbers[1:]) ||
		rec2(concat, target, numbers[1:]) {
		return true
	}

	return false
}

func part1and2(f *os.File) (int, int) {
	r1, r2 := 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		numbers := make([]int, 0)
		for _, numberStr := range strings.Split(parts[1], " ") {
			number, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, number)
		}
		if rec(0, target, numbers, false) {
			r1 += target
		}
		if rec(0, target, numbers, true) {
			r2 += target
		}
	}
	return r1, r2
}

func rec(start, target int, numbers []int, isPart2 bool) bool {
	if len(numbers) == 0 && start == target {
		return true
	}
	if len(numbers) == 0 {
		return false
	}
	if rec(start+numbers[0], target, numbers[1:], isPart2) ||
		rec(start*numbers[0], target, numbers[1:], isPart2) {
		return true
	}
	if isPart2 {
		concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", start, numbers[0]))
		if rec(concat, target, numbers[1:], isPart2) {
			return true
		}
	}
	return false
}
