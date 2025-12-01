package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f1, _ := os.Open("2025/day1/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2025/day1/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0
	value := 50

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		turns, _ := strconv.Atoi(line[1:])

		switch direction {
		case 'L':
			v := value - turns
			for v < 0 {
				v += 100
			}
			value = v
		case 'R':
			v := value + turns
			if v%100 == 0 {
				v = 0
			}
			for v > 100 {
				v -= 100
			}
			value = v
		}

		if value == 0 {
			result += 1
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0
	value := 50

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		turns, _ := strconv.Atoi(line[1:])
		result += turns / 100
		turns = turns % 100

		switch direction {
		case 'L':
			value -= turns
			if value == 0 {
				result += 1
			} else if value < 0 {
				value += 100
				if value+turns != 100 {
					result += 1
				}
			}
		case 'R':
			value += turns
			if value > 99 {
				value -= 100
				if value-turns != 0 {
					result += 1
				}
			} else if value == 0 {
				result += 1
			}
		}
	}

	return result
}
