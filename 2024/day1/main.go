package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f1, _ := os.Open("2024/day1/input.txt")
	defer f1.Close()
	fmt.Println("Part 1:", part1(f1))

	f2, _ := os.Open("2024/day1/input.txt")
	defer f2.Close()
	fmt.Println("Part 2:", part2(f2))
}

func part1(f *os.File) int {
	result := 0

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())

		numInt1, _ := strconv.Atoi(nums[0])
		numInt2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, numInt1)
		list2 = append(list2, numInt2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := range list1 {
		result += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return result
}

func part2(f *os.File) int {
	result := 0

	list := make([]int, 0)
	m := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())

		num1, _ := strconv.Atoi(nums[0])
		list = append(list, num1)

		num2, _ := strconv.Atoi(nums[1])
		if _, ok := m[num2]; ok {
			m[num2]++
		} else {
			m[num2] = 1
		}
	}

	for _, v := range list {
		if _, ok := m[v]; ok {
			result += v * m[v]
		}
	}

	return result
}
