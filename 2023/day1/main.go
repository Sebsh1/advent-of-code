package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2023/day1/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2023/day1/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

var wordMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part1(f *os.File) int {
	result := 0

	regex := regexp.MustCompile(`\d`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := regex.FindAllString(scanner.Text(), -1)

		num1, _ := strconv.Atoi(string(matches[0]))
		num2, _ := strconv.Atoi(string(matches[len(matches)-1]))

		result += num1*10 + num2
	}

	return result
}

func swapWordForDigit(s string) (bool, string) {
	for k, v := range wordMap {
		if strings.Contains(s, k) {
			return true, strings.Replace(s, k, strconv.Itoa(v), 1)
		}
	}
	return false, s
}

func part2(f *os.File) int {
	result := 0

	regex := regexp.MustCompile(`\d`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		finished := false
		for !finished {
			for i := range line {
				if found, newLine := swapWordForDigit(line[:i]); found {
					line = newLine + line[i:]
					break
				}

				if i == len(line)-1 {
					finished = true
				}
			}
		}

		matches := regex.FindAllString(line, -1)

		num1, _ := strconv.Atoi(string(matches[0]))
		num2, _ := strconv.Atoi(string(matches[len(matches)-1]))

		result += num1*10 + num2
	}

	return result
}
