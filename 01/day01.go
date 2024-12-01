package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/osagemo/advent-of-code-24/internal/utils"
)

//go:embed input.txt
var input string

// Pair up the smallest number in the left list with the smallest number in the right list,
// then the second-smallest left number with the second-smallest right number, and so on.
// Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances.
func Part1(input string) int {
	totalDistance := 0

	leftList := []int{}
	rightList := []int{}

	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		split := strings.Fields(line)
		leftList = append(leftList, utils.MustParseInt(split[0]))
		rightList = append(rightList, utils.MustParseInt(split[1]))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	for i := 0; i < len(leftList); i++ {
		totalDistance += utils.Abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

// Calculate a total similarity score by adding up each number in the left list
// after multiplying it by the number of times that number appears in the right list.
func Part2(input string) int {
	totalScore := 0

	leftList := []int{}
	rightCounts := map[int]int{}

	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		split := strings.Fields(line)
		leftList = append(leftList, utils.MustParseInt(split[0]))
		if num, err := strconv.Atoi(split[1]); err == nil {
			if _, ok := rightCounts[num]; !ok {
				rightCounts[num] = 1
			} else {
				rightCounts[num]++
			}
		}
	}

	for i := 0; i < len(leftList); i++ {
		num := leftList[i]
		totalScore += num * rightCounts[num]
	}

	return totalScore
}

func main() {
	input := strings.ReplaceAll(input, "\r\n", "\n")
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}
