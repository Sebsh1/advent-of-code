package main

import (
	"fmt"
	"os"
)

func main() {
	f1, _ := os.Open("2024/day4/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2024/day4/input.txt")
	defer f2.Close()
	fmt.Println("Part 2: ", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	return result
}

func part2(f *os.File) int {
	result := 0

	return result
}
