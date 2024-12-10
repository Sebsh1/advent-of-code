package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day10/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	start = time.Now()
	result = part2()
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

type point struct {
	x, y int
}

var trailheads = make([]point, 0)
var topographicMap = make(map[point]int, 0)
var reachablePoints = make(map[point][]point, 0)
var uniqueTrails = make(map[point][]point, 0)

func part1(f *os.File) int {
	y := -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		y++
		line := scanner.Text()
		for x, c := range line {
			height := int(c - '0')
			topographicMap[point{x, y}] = height
			if height == 0 {
				trailheads = append(trailheads, point{x, y})
			}
		}
	}

	for _, th := range trailheads {
		rec1(th, th)
	}

	result := 0
	for _, v := range reachablePoints {
		result += len(v)
	}

	return result
}

func part2() int {
	for _, th := range trailheads {
		rec2(th, th)
	}

	result := 0
	for _, v := range uniqueTrails {
		result += len(v)
	}

	return result
}

func getNeighbors(p point) []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

func rec1(start, current point) {
	for _, n := range getNeighbors(current) {
		height, ok := topographicMap[n]
		if ok && height == 9 && topographicMap[current] == 8 && !slices.Contains(reachablePoints[start], n) {
			reachablePoints[start] = append(reachablePoints[start], n)
		} else if ok && height == topographicMap[current]+1 {
			rec1(start, n)
		}
	}
}

func rec2(start, current point) {
	for _, n := range getNeighbors(current) {
		height, ok := topographicMap[n]
		if ok && height == 9 && topographicMap[current] == 8 {
			uniqueTrails[start] = append(uniqueTrails[start], n)
		} else if ok && height == topographicMap[current]+1 {
			rec2(start, n)
		}
	}
}
