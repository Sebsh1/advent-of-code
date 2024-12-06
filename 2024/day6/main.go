package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day6/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	f2, _ := os.Open("2024/day6/input.txt")
	defer f2.Close()
	start = time.Now()
	result = part2(f2)
	end = time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)
}

type position struct {
	x int
	y int
}

type direction struct {
	x int
	y int
}

var guardSymbols = []rune{'^', '>', 'v', '<'}

func part1(f *os.File) int {
	result := 0
	out := false
	grid := make([][]rune, 0)
	guardPosition := position{0, 0}
	guardDirection := direction{0, 0}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		for i, letter := range line {
			if slices.Contains(guardSymbols, letter) {
				guardPosition = position{i, len(grid) - 1}
				guardDirection = guardDirFromSymbol(letter)
			}
		}
	}

	for !out {
		nextPosition := position{guardPosition.x + guardDirection.x, guardPosition.y + guardDirection.y}
		if nextPosition.x < 0 || nextPosition.x >= len(grid[guardPosition.y]) || nextPosition.y < 0 || nextPosition.y >= len(grid) {
			out = true
			result++
			break
		}
		if grid[nextPosition.y][nextPosition.x] == '#' {
			guardDirection = rotate(guardDirection)
			nextPosition = position{guardPosition.x + guardDirection.x, guardPosition.y + guardDirection.y}
		}
		if grid[guardPosition.y][guardPosition.x] != 'X' {
			result++
		}
		grid[guardPosition.y][guardPosition.x] = 'X'
		guardPosition = nextPosition
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	grid := make([][]rune, 0)
	startGuardPosition := position{0, 0}
	startGuardDirection := direction{0, 0}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		for i, letter := range line {
			if slices.Contains(guardSymbols, letter) {
				startGuardPosition = position{i, len(grid) - 1}
				startGuardDirection = guardDirFromSymbol(letter)
			}
		}
	}

	checks := 0
	for i, line := range grid {
		for j := range line {
			guardPosition := startGuardPosition
			guardDirection := startGuardDirection

			if guardPosition.x == j && guardPosition.y == i {
				continue
			}
			checks++

			cellVisits := make(map[position]int)
			gridCopy := copyGrid(grid)
			gridCopy[i][j] = '#'
			out := false
			loop := false
			for !out && !loop {
				cellVisits[guardPosition]++
				nextPosition := position{guardPosition.x + guardDirection.x, guardPosition.y + guardDirection.y}
				if nextPosition.x < 0 || nextPosition.x >= len(gridCopy[guardPosition.y]) || nextPosition.y < 0 || nextPosition.y >= len(gridCopy) {
					out = true
					break
				}
				if gridCopy[nextPosition.y][nextPosition.x] == '#' {
					guardDirection = rotate(guardDirection)
					nextPosition = position{guardPosition.x + guardDirection.x, guardPosition.y + guardDirection.y}
				}
				guardPosition = nextPosition

				if cellVisits[guardPosition] > 4 {
					loop = true
					result++
				}
			}
		}
	}

	return result
}

func guardDirFromSymbol(s rune) direction {
	switch s {
	case '^':
		return direction{0, -1}
	case '>':
		return direction{1, 0}
	case 'v':
		return direction{0, 1}
	case '<':
		return direction{-1, 0}
	}
	return direction{0, 0}
}

func rotate(dir direction) direction {
	switch dir {
	case direction{-1, 0}:
		return direction{0, -1}
	case direction{0, -1}:
		return direction{1, 0}
	case direction{1, 0}:
		return direction{0, 1}
	case direction{0, 1}:
		return direction{-1, 0}
	}
	return direction{0, 0}
}

func printGrid(grid [][]rune) {
	fmt.Println()
	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func copyGrid(grid [][]rune) [][]rune {
	gridCopy := make([][]rune, len(grid))
	for i, line := range grid {
		gridCopy[i] = make([]rune, len(line))
		copy(gridCopy[i], line)
	}
	return gridCopy
}
