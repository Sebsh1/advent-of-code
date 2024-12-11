package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f1, _ := os.Open("2022/day1/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2022/day1/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	calories := make(map[int]int, 0)
	elf := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elf++
		} else {
			cal, _ := strconv.Atoi(line)
			calories[elf] += cal
		}
	}

	maxCal := 0
	for _, v := range calories {
		if v > maxCal {
			maxCal = v
		}
	}

	return maxCal
}

func part2(f *os.File) int {
	calories := make(map[int]int, 0)
	elf := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elf++
		} else {
			cal, _ := strconv.Atoi(line)
			calories[elf] += cal
		}
	}

	firstCal, secondCal, thirdCal := 0, 0, 0
	for _, v := range calories {
		if v > firstCal {
			thirdCal = secondCal
			secondCal = firstCal
			firstCal = v
		} else if v > secondCal {
			thirdCal = secondCal
			secondCal = v
		} else if v > thirdCal {
			thirdCal = v
		}
	}

	return firstCal + secondCal + thirdCal
}
