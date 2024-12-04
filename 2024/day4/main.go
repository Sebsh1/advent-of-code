package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f1, _ := os.Open("2024/day4/input.txt")
	defer f1.Close()
	fmt.Println("Part 1: ", part1(f1))

	matrix = nil

	f2, _ := os.Open("2024/day4/input.txt")
	defer f2.Close()
	fmt.Println("Part 2: ", part2(f2))
}

var matrix []string

var diagonals = [][]int{
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

var directions = append(diagonals, [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}...)

func part1(f *os.File) int {
	result := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	for i, row := range matrix {
		for j, letter := range row {
			if letter == 'X' {
				dirs := letterAdjacentDirections(i, j, 'M', directions)
				for _, dir := range dirs {
					if isDirectionXMAS(i+dir[0], j+dir[1], dir) {
						result++
					}
				}
			}
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	for i, row := range matrix {
		for j, letter := range row {
			if letter == 'A' {
				mDirs := letterAdjacentDirections(i, j, 'M', diagonals)
				sDirs := letterAdjacentDirections(i, j, 'S', diagonals)

				if len(mDirs) != 2 || len(sDirs) != 2 {
					continue
				}

				if matrix[i-1][j-1] == matrix[i+1][j+1] {
					continue
				}

				result++
			}
		}
	}

	return result
}

func letterAdjacentDirections(i, j int, letter rune, directions [][]int) [][]int {
	dirs := make([][]int, 0)
	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && rune(matrix[x][y]) == letter {
			dirs = append(dirs, dir)
		}
	}

	return dirs
}

func isDirectionXMAS(i, j int, dir []int) bool {
	x, y := i+dir[0], j+dir[1]
	if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == 'A' {
		x, y = x+dir[0], y+dir[1]
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == 'S' {
			return true
		}
	}

	return false
}
