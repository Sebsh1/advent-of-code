package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}

func main() {
	f1, _ := os.Open("2023/day2/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2023/day2/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		g := parseGame(scanner.Text())

		valid := true
		for _, s := range g.sets {
			if s.red > 12 || s.green > 13 || s.blue > 14 {
				valid = false
				break
			}
		}

		if valid {
			result += g.id
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		g := parseGame(scanner.Text())

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, s := range g.sets {
			if s.red > maxRed {
				maxRed = s.red
			}
			if s.green > maxGreen {
				maxGreen = s.green
			}
			if s.blue > maxBlue {
				maxBlue = s.blue
			}
		}

		result += maxRed * maxGreen * maxBlue
	}

	return result
}

func parseGame(line string) game {
	parts := strings.Split(line, ": ")

	gameIDStr := strings.TrimPrefix(parts[0], "Game ")
	gameID, _ := strconv.Atoi(gameIDStr)

	g := game{id: gameID}

	setsStr := parts[1]
	setStrings := strings.Split(setsStr, "; ")
	for _, setStr := range setStrings {
		s := set{}

		cubes := strings.Split(setStr, ", ")
		for _, cubeStr := range cubes {
			parts := strings.Split(cubeStr, " ")
			count, _ := strconv.Atoi(parts[0])

			switch parts[1] {
			case "red":
				s.red = count
			case "green":
				s.green = count
			case "blue":
				s.blue = count
			}
		}

		g.sets = append(g.sets, s)
	}

	return g
}
