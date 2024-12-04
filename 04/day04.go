package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	matrix := parseMatrix(input)
	return countWordOccurrences(matrix, "XMAS")
}

func Part2(input string) int {
	matrix := parseMatrix(input)
	return countCrossmasOccurrences(matrix)
}

var allDirections = []Coordinate{
	{0, 1}, {1, 0}, {1, 1},
	{0, -1}, {-1, 0}, {-1, -1},
	{-1, 1}, {1, -1},
}

func countWordOccurrences(matrix [][]rune, word string) int {
	count := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == rune(word[0]) {
				count += searchWordInAllDirections(matrix, word, Coordinate{Row: row, Col: col})
			}
		}
	}
	return count
}

func searchWordInAllDirections(matrix [][]rune, word string, start Coordinate) int {
	matchCount := 0
	for _, dir := range allDirections {
		if wordInDirection(matrix, word, dir, start) {
			matchCount++
		}
	}
	return matchCount
}

func wordInDirection(matrix [][]rune, word string, direction Coordinate, start Coordinate) bool {
	position := start
	for _, char := range word {
		if !withinBounds(matrix, position) || matrix[position.Row][position.Col] != char {
			return false
		}
		position = position.Add(direction)
	}
	return true
}

// it's an X-MAS puzzle in which you're supposed to find two MAS
// in the shape of an X. One way to achieve that is like this:
// M.S
// .A.
// M.S
// (lol)
func countCrossmasOccurrences(matrix [][]rune) int {
	count := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 'A' {
				if isCrossmasFound(matrix, Coordinate{Row: row, Col: col}) {
					count++
				}
			}
		}
	}
	return count
}

func isCrossmasFound(matrix [][]rune, start Coordinate) bool {
	firstDiagonal := []Coordinate{{1, 1}, {-1, -1}}
	secondDiagonal := []Coordinate{{-1, 1}, {1, -1}}

	return isDiagonalValid(matrix, start, firstDiagonal) &&
		isDiagonalValid(matrix, start, secondDiagonal)
}

func isDiagonalValid(matrix [][]rune, start Coordinate, directions []Coordinate) bool {
	foundM, foundS := false, false
	for _, direction := range directions {
		pos := start.Add(direction)
		if withinBounds(matrix, pos) {
			switch matrix[pos.Row][pos.Col] {
			case 'M':
				foundM = true
			case 'S':
				foundS = true
			}
		}
	}
	return foundM && foundS
}

func withinBounds(matrix [][]rune, coord Coordinate) bool {
	return coord.Row >= 0 && coord.Row < len(matrix) &&
		coord.Col >= 0 && coord.Col < len(matrix[0])
}

func parseMatrix(input string) [][]rune {
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
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

type Coordinate struct {
	Row, Col int
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{Row: c.Row + other.Row, Col: c.Col + other.Col}
}
