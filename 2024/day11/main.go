package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day11/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	f2, _ := os.Open("2024/day11/input.txt")
	defer f1.Close()
	start = time.Now()
	result = part2(f2)
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

func part1(f *os.File) int {
	stones := make([]int, 0)
	content, _ := io.ReadAll(f)
	for _, numStr := range strings.Split(string(content), " ") {
		stone, _ := strconv.Atoi(numStr)
		stones = append(stones, processStone(stone, 25)...)
	}

	return len(stones)
}

func part2(f *os.File) int {
	evens := make(map[int]int, 0)
	odds := make(map[int]int, 0)

	content, _ := io.ReadAll(f)
	for _, numStr := range strings.Split(string(content), " ") {
		stone, _ := strconv.Atoi(numStr)
		if len(strconv.Itoa(stone))%2 == 0 {
			evens[stone]++
		} else {
			odds[stone]++
		}
	}

	for range 75 {
		evens, odds = processStones(evens, odds)
	}

	result := 0
	for _, count := range evens {
		result += count
	}
	for _, count := range odds {
		result += count
	}

	return result
}

func processStone(stone int, blinks int) []int {
	stones := []int{stone}
	for range blinks {
		updatedStones := make([]int, 0, len(stones))
		for _, s := range stones {
			if s == 0 {
				updatedStones = append(updatedStones, 1)
			} else if sStr := strconv.Itoa(s); len(sStr)%2 == 0 {
				s1, _ := strconv.Atoi(sStr[:len(sStr)/2])
				s2, _ := strconv.Atoi(sStr[len(sStr)/2:])
				updatedStones = append(updatedStones, s1, s2)
			} else {
				updatedStones = append(updatedStones, s*2024)
			}
		}
		stones = updatedStones
	}
	return stones
}

func processStones(evens, odds map[int]int) (map[int]int, map[int]int) {
	newEvens := make(map[int]int, 0)
	newOdds := make(map[int]int, 0)

	for num, count := range evens {
		numStr := strconv.Itoa(num)
		num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
		if len(strconv.Itoa(num1))%2 == 0 {
			newEvens[num1] += count
		} else {
			newOdds[num1] += count
		}
		num2, _ := strconv.Atoi(numStr[len(numStr)/2:])
		if len(strconv.Itoa(num2))%2 == 0 {
			newEvens[num2] += count
		} else {
			newOdds[num2] += count
		}
	}

	for num, count := range odds {
		if num == 0 {
			newOdds[1] += count
		} else {
			newNum := num * 2024
			if len(strconv.Itoa(newNum))%2 == 0 {
				newEvens[newNum] += count
			} else {
				newOdds[newNum] += count
			}
		}
	}

	return newEvens, newOdds
}
