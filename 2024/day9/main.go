package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	f1, _ := os.Open("2024/day9/input.txt")
	defer f1.Close()
	start := time.Now()
	result := part1(f1)
	end := time.Since(start)
	fmt.Println("Part 1:", result, "Time:", end)

	f2, _ := os.Open("2024/day9/input.txt")
	defer f1.Close()
	start = time.Now()
	result = part2(f2)
	end = time.Since(start)
	fmt.Println("Part 2:", result, "Time:", end)
}

func part1(f *os.File) int {
	result := 0
	disk := make([]rune, 0)

	lineBytes, _ := io.ReadAll(f)
	line := string(lineBytes)
	for i, r := range line {
		if i%2 == 0 {
			for j := 0; j < int(r-'0'); j++ {
				disk = append(disk, rune('0'+i/2))
			}
		} else {
			for j := 0; j < int(r-'0'); j++ {
				disk = append(disk, '.')
			}
		}
	}

	for i, j := 0, len(disk)-1; i <= j; {
		if disk[i] == '.' && disk[j] != '.' {
			disk[i], disk[j] = disk[j], disk[i]
			i++
			j--
		} else {
			if disk[i] != '.' {
				i++
			} else if disk[j] != '.' {
				j--
			} else if disk[i] == '.' && disk[j] == '.' {
				j--
			}
		}
	}

	for i := range disk {
		if disk[i] != '.' {
			addToResult := i * int(disk[i]-'0')
			result += addToResult
		}
	}

	return result
}

func part2(f *os.File) int {
	result := 0
	disk := make([]rune, 0)
	files := make([][]int, 0)
	spaces := make([][]int, 0)

	lineBytes, _ := io.ReadAll(f)
	line := string(lineBytes)
	for i, r := range line {
		if i%2 == 0 {
			files = append(files, []int{len(disk), len(disk) - 1 + int(r-'0')})
			for j := 0; j < int(r-'0'); j++ {
				disk = append(disk, rune('0'+i/2))
			}
		} else {
			spaces = append(spaces, []int{len(disk), len(disk) - 1 + int(r-'0')})
			for j := 0; j < int(r-'0'); j++ {
				disk = append(disk, '.')
			}
		}
	}

	for j := len(files) - 1; j >= 0; j-- {
		for _, space := range spaces {
			if space[1] < files[j][1] && space[1]-space[0] >= files[j][1]-files[j][0] {
				for k := space[0]; k <= space[0]+files[j][1]-files[j][0]; k++ {
					disk[k] = rune('0' + j)
				}
				for k := files[j][0]; k <= files[j][1]; k++ {
					disk[k] = '.'
				}
				space[0] = space[0] + files[j][1] - files[j][0] + 1
				break
			}
		}

	}

	for i := range disk {
		if disk[i] != '.' {
			result += i * int(disk[i]-'0')
		}
	}

	return result
}
