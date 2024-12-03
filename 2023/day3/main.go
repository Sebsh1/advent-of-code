package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f1, _ := os.Open("2023/day3/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2023/day3/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

type number struct {
	value      string
	lineIndex  int
	startIndex int
	endIndex   int
}

func part1(f *os.File) int {
	result := 0
	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Identify all the numbers
	numbers := make([]number, 0)
	for i, line := range lines {
		isNumber := false
		num := ""
		startIndex := 0
		for j, char := range line {
			if unicode.IsDigit(char) {
				if !isNumber {
					startIndex = j
				}
				isNumber = true
				num += string(char)
			} else if isNumber {
				numbers = append(numbers, number{
					value:      num,
					lineIndex:  i,
					startIndex: startIndex,
					endIndex:   j,
				})
				num = ""
				isNumber = false
			}
		}
		if isNumber { // Handle number at end of line
			numbers = append(numbers, number{
				value:      num,
				lineIndex:  i,
				startIndex: startIndex,
				endIndex:   len(line),
			})
		}
	}

	// Process numbers
	for _, num := range numbers {
		if ok, _, _ := isAdjacentToCriteria(num, lines, isSymbol); ok {
			numInt, _ := strconv.Atoi(num.value)
			result += numInt
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	type symbol struct {
		x int
		y int
	}

	symbolAdjacentCount := make(map[symbol][]int)

	// Identify all the numbers
	numbers := make([]number, 0)
	for i, line := range lines {
		isNumber := false
		num := ""
		startIndex := 0
		for j, char := range line {
			if unicode.IsDigit(char) {
				if !isNumber {
					startIndex = j
				}
				isNumber = true
				num += string(char)
			} else if isNumber {
				numbers = append(numbers, number{
					value:      num,
					lineIndex:  i,
					startIndex: startIndex,
					endIndex:   j,
				})
				num = ""
				isNumber = false
			}
		}
		if isNumber { // Handle number at end of line
			numbers = append(numbers, number{
				value:      num,
				lineIndex:  i,
				startIndex: startIndex,
				endIndex:   len(line),
			})
		}
	}

	// Process numbers
	for _, num := range numbers {
		if ok, x, y := isAdjacentToCriteria(num, lines, isGear); ok {
			numInt, _ := strconv.Atoi(num.value)
			symbolAdjacentCount[symbol{x, y}] = append(symbolAdjacentCount[symbol{x, y}], numInt)
		}
	}

	// Find gear ratios
	for _, nums := range symbolAdjacentCount {
		if len(nums) != 2 {
			continue
		}

		result += nums[0] * nums[1]
	}

	return result
}

// Returns if the number is adjacent to a criteria and the x,y indeces of the symbol
func isAdjacentToCriteria(num number, lines []string, criteria func(char byte) bool) (bool, int, int) {
	startSearchIndex := max(0, num.startIndex-1)
	endSearchIndex := min(len(lines[num.lineIndex]), num.endIndex+1)

	// Check same line before and after number
	if num.startIndex > 0 && criteria(lines[num.lineIndex][num.startIndex-1]) {
		return true, num.startIndex - 1, num.lineIndex
	}
	if num.endIndex < len(lines[num.lineIndex]) && criteria(lines[num.lineIndex][num.endIndex]) {
		return true, num.endIndex, num.lineIndex
	}

	// Check lines above and below
	if num.lineIndex > 0 {
		if s, i := checkLineForCriteria(lines[num.lineIndex-1], startSearchIndex, endSearchIndex, criteria); s {
			return true, i, num.lineIndex - 1
		}
	}
	if num.lineIndex < len(lines)-1 {
		if s, i := checkLineForCriteria(lines[num.lineIndex+1], startSearchIndex, endSearchIndex, criteria); s {
			return true, i, num.lineIndex + 1
		}
	}

	return false, -1, -1
}

// Returns if the string contains a symbol and the index of the symbol
func checkLineForCriteria(line string, start, end int, criteria func(char byte) bool) (bool, int) {
	for i := start; i < end; i++ {
		if criteria(line[i]) {
			return true, i
		}
	}
	return false, -1
}

func isSymbol(char byte) bool {
	return !unicode.IsDigit(rune(char)) && char != '.'
}

func isGear(char byte) bool {
	return char == '*'
}
