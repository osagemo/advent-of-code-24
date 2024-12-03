package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/osagemo/advent-of-code-24/internal/utils"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	instructions := ParseInstructions(input)
	sum := 0
	for _, i := range instructions {
		sum += ExecuteInstruction(i)
	}

	return sum
}

func Part2(input string) int {
	instructions := ParseInstructions2(input)
	sum := 0
	for _, i := range instructions {
		sum += ExecuteInstruction(i)
	}

	return sum
}

func ParseInstructions(input string) []Instruction {
	instructions := []Instruction{}
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		instructions = append(instructions, Instruction{
			X:         utils.MustParseInt(match[1]),
			Y:         utils.MustParseInt(match[2]),
			Operation: Mul,
		})
	}

	return instructions
}

func ParseInstructions2(input string) []Instruction {
	const (
		Enable  = "do()"
		Disable = "don't()"
	)
	instructions := []Instruction{}
	re := regexp.MustCompile(`do\(\)|mul\((\d+),(\d+)\)|don't\(\)`)
	matches := re.FindAllString(input, -1)

	enabled := true
	for _, match := range matches {
		if match == Enable {
			enabled = true
			continue
		}
		if match == Disable {
			enabled = false
			continue
		}
		if enabled {
			re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
			instructionMatch := re.FindStringSubmatch(match)
			instructions = append(instructions, Instruction{
				X:         utils.MustParseInt(instructionMatch[1]),
				Y:         utils.MustParseInt(instructionMatch[2]),
				Operation: Mul,
			})
		}
	}

	return instructions
}

func ExecuteInstruction(i Instruction) int {
	return i.X * i.Y
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

type Operation string

const (
	Mul Operation = "mul"
)

type Instruction struct {
	X         int
	Y         int
	Operation Operation
}
