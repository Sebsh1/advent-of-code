package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2025/day2/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2025/day2/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	intervals := strings.SplitSeq(line, ",")

	for interval := range intervals {
		parts := strings.Split(interval, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			num := strconv.Itoa(i)
			if len(num)%2 != 0 {
				continue
			}

			mid := len(num) / 2
			if num[:mid] == num[mid:] {
				result += i
			}
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	intervals := strings.SplitSeq(line, ",")

	for interval := range intervals {
		parts := strings.Split(interval, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			num := strconv.Itoa(i)
			mid := len(num) / 2

			for j := 1; j <= mid; j++ {
				if len(num)%j != 0 {
					continue
				}
				if strings.Count(num, num[:j]) == len(num)/j {
					result += i
					break
				}
			}
		}
	}

	return result
}
