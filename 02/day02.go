package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/osagemo/advent-of-code-24/internal/utils"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	totalSafeReports := 0
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		levels := utils.MustParseInts(strings.Fields(line))
		if isSafe(levels) {
			totalSafeReports++
		}
	}

	return totalSafeReports
}

func Part2(input string) int {
	totalSafeReports := 0
	for _, line := range strings.Split(input, "\n") {
		levels := utils.MustParseInts(strings.Fields(line))
		if isSafe(levels) {
			totalSafeReports++
		} else {
			// Brute force, remove one level at a time and check if it's safe
			for i := 0; i < len(levels); i++ {
				if isSafe(utils.RemoveIndex(levels, i)) {
					totalSafeReports++
					break
				}
			}
		}
	}

	return totalSafeReports
}

func isSafe(levels []int) bool {
	prev := levels[0]
	increasing := levels[1] > levels[0]
	safe := true
	for i, n := range levels {
		if i == 0 {
			continue
		}

		dif := n - prev
		// Check sign?
		if dif >= 0 && !increasing {
			safe = false
			break
		}
		if dif <= 0 && increasing {
			safe = false
			break
		}
		if utils.Abs(dif) > 3 {
			safe = false
			break
		}

		prev = n
	}

	return safe
}

func main() {
	input := strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.Trim(input, "\n")
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}
