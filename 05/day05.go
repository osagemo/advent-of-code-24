package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/osagemo/advent-of-code-24/internal/utils"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	totals := 0
	sorted, _ := separateSortedAndUnsortedUpdates(input)

	for _, updates := range sorted {
		totals += getMiddle(updates)
	}

	return totals
}

func Part2(input string) int {
	totals := 0

	pageOrderingRules, _ := splitInput(input)
	smallerToBigger := mapSmallerToBigger(pageOrderingRules)
	_, unsorted := separateSortedAndUnsortedUpdates(input)

	for _, updates := range unsorted {
		slices.SortFunc(updates, func(a, b int) int {
			biggerThan := smallerToBigger[a]
			if slices.Contains(biggerThan, b) {
				return 1
			}
			return -1
		})
		totals += getMiddle(updates)
	}

	return totals
}

func getMiddle(s []int) int {
	if len(s) < 2 {
		panic("no mid")
	}
	return s[len(s)/2]
}

func separateSortedAndUnsortedUpdates(input string) (sorted [][]int, unsorted [][]int) {
	pageOrderingRules, updates := splitInput(input)
	smallerToBigger := mapSmallerToBigger(pageOrderingRules)
	for _, line := range strings.Split(updates, "\n") {
		pageNumbers := utils.MustParseInts(strings.Split(line, ","))
		updateIsSorted := true
		for i := 0; i < len(pageNumbers)-1; i++ {
			biggerThan := smallerToBigger[pageNumbers[i]]
			if !slices.Contains(biggerThan, pageNumbers[i+1]) {
				updateIsSorted = false
				break
			}
		}

		if updateIsSorted {
			sorted = append(sorted, pageNumbers)
		} else {
			unsorted = append(unsorted, pageNumbers)
		}
	}
	return
}

func splitInput(input string) (pageOrderingRules string, updates string) {
	sections := strings.Split(input, "\n\n")
	pageOrderingRules, updates = sections[0], sections[1]
	return
}

func mapSmallerToBigger(pageOrderingRulesData string) map[int][]int {
	smallerThanMap := map[int][]int{}
	for _, line := range strings.Split(pageOrderingRulesData, "\n") {
		split := strings.Split(line, "|")
		smaller, bigger := utils.MustParseInt(split[0]), utils.MustParseInt(split[1])
		smallerThanMap[smaller] = append(smallerThanMap[smaller], bigger)
	}

	return smallerThanMap
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
