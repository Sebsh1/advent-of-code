package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f1, _ := os.Open("2025/day4/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2025/day4/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

var directions = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func getAdjacent(grid []string, x, y int) []byte {
	adjacent := make([]byte, 0)
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) {
			adjacent = append(adjacent, grid[newY][newX])
		}
	}

	return adjacent
}

func part1(f *os.File) int {
	grid := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	result := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '@' {
				continue
			}
			count := 0
			for _, adj := range getAdjacent(grid, x, y) {
				if adj == '@' {
					count++
				}
			}
			if count < 4 {
				result++
			}
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0
	toRemove := make(map[[2]int]struct{})
	grid := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

loop:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '@' {
				continue
			}
			count := 0
			for _, adj := range getAdjacent(grid, x, y) {
				if adj == '@' {
					count++
				}
			}
			if count < 4 {
				result++
				toRemove[[2]int{x, y}] = struct{}{}
			}
		}
	}

	for pos := range toRemove {
		grid[pos[1]] = grid[pos[1]][:pos[0]] + "." + grid[pos[1]][pos[0]+1:]
	}
	if len(toRemove) > 0 {
		toRemove = make(map[[2]int]struct{})
		goto loop
	}

	return result
}
