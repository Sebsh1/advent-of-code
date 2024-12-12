package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day12/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	f2, _ := os.Open("2024/day12/input.txt")
	defer f1.Close()
	start = time.Now()
	result = part2(f2)
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

type coord struct {
	x, y int
}

var dirs = []coord{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

var grid = make([][]rune, 0)
var plots = make(map[int][]coord, 0)
var partOfPlot = make(map[coord]bool, 0)

func part1(f *os.File) int {
	result := 0
	id := 0

	content, _ := io.ReadAll(f)
	for _, line := range strings.Split(string(content), "\n") {
		grid = append(grid, []rune(line))
	}

	for y := range grid {
		for x := range grid[y] {
			id++
			c := coord{x, y}
			if !partOfPlot[c] {
				plots[id] = []coord{c}
				partOfPlot[c] = true
			}
			findAdjacent(c, id, grid)
		}
	}

	for _, plot := range plots {
		fences := 0
		for _, p := range plot {
			fences += requiredFences(p)
		}

		result += fences * len(plot)
	}

	return result
}

func findAdjacent(c coord, id int, grid [][]rune) {
	for _, dir := range dirs {
		n := coord{c.x + dir.x, c.y + dir.y}
		if n.x < 0 || n.y < 0 || n.y > len(grid)-1 || n.x > len(grid[c.y])-1 {
			continue
		}
		if !partOfPlot[n] && grid[n.y][n.x] == grid[c.y][c.x] {
			plots[id] = append(plots[id], n)
			partOfPlot[n] = true
			findAdjacent(n, id, grid)
		}
	}
}

func requiredFences(c coord) int {
	fences := 0
	for _, dir := range dirs {
		n := coord{c.x + dir.x, c.y + dir.y}
		if n.x < 0 || n.y < 0 || n.y > len(grid)-1 || n.x > len(grid[c.y])-1 {
			fences++
		} else if grid[n.y][n.x] != grid[c.y][c.x] {
			fences++
		}
	}
	return fences
}

func part2(f *os.File) int {
	result := 0

	return result
}
