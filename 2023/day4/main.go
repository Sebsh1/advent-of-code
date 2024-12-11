package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2023/day4/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2023/day4/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	regex := regexp.MustCompile(`\d{1,2}`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		leftParts := strings.Split(parts[0], ":")
		numbers := regex.FindAllString(leftParts[1], -1)
		winningNumbers := regex.FindAllString(parts[1], -1)

		score := 0
		for _, num := range numbers {
			if slices.Contains(winningNumbers, num) {
				if score == 0 {
					score += 1
				} else {
					score *= 2
				}
			}
		}

		result += score
	}

	return result
}

// TODO - implement part2?
func part2(f *os.File) int {
	result := 0

	regex := regexp.MustCompile(`\d{1,3}`)

	cardCounts := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		leftParts := strings.Split(parts[0], ":")

		cardNumbers := regex.FindAllString(leftParts[0], 1)
		cardNumber, _ := strconv.Atoi(cardNumbers[0])

		numbers := regex.FindAllString(leftParts[1], -1)
		winningNumbers := regex.FindAllString(parts[1], -1)

		count := 0
		for _, num := range numbers {
			if slices.Contains(winningNumbers, num) {
				count++
			}
		}

		for range cardCounts[cardNumber] {
			for j := range count {
				cardCounts[cardNumber+1+j] = cardCounts[cardNumber+1+j] + 1
			}
		}
	}

	for _, v := range cardCounts {
		result += v
	}

	return result
}
