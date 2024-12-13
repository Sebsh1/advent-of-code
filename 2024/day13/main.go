package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day13/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	f2, _ := os.Open("2024/day13/input.txt")
	defer f1.Close()
	start = time.Now()
	result = part2(f2)
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

type machine struct {
	aX, aY         int
	bX, bY         int
	prizeX, prizeY int
}

func part1(f *os.File) int {
	result := 0
	content, _ := io.ReadAll(f)
	for _, line := range strings.Split(string(content), "\n\n") {
		m := machine{}
		fmt.Sscanf(line, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &m.aX, &m.aY, &m.bX, &m.bY, &m.prizeX, &m.prizeY)
		result += pressButtons(m)
	}
	return result
}

func pressButtons(m machine) int {
	for a := range 101 {
		for b := range 101 {
			if m.aX*a+m.bX*b == m.prizeX && m.aY*a+m.bY*b == m.prizeY {
				return a*3 + b
			}
		}
	}
	return 0
}

func part2(f *os.File) int {
	result := 0
	content, _ := io.ReadAll(f)
	for _, line := range strings.Split(string(content), "\n\n") {
		m := machine{}
		fmt.Sscanf(line, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &m.aX, &m.aY, &m.bX, &m.bY, &m.prizeX, &m.prizeY)
		m.prizeX = m.prizeX + 10000000000000
		m.prizeY = m.prizeY + 10000000000000
		result += solveEquations(m)
	}
	return result
}

func solveEquations(m machine) int {
	d := m.aX*m.bY - m.aY*m.bX
	if d == 0 {
		return 0
	}
	a := float64((m.prizeX*m.bY - m.prizeY*m.bX)) / float64(d)
	b := float64(m.aX*m.prizeY-m.aY*m.prizeX) / float64(d)
	if a == float64(int(a)) && b == float64(int(b)) {
		return int(a)*3 + int(b)
	}
	return 0
}
