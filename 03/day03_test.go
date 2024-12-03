package main

import (
	"fmt"
	"testing"
)

const input1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const input2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestParseInstructions(t *testing.T) {
	instructions := ParseInstructions(input1)
	expected := []Instruction{
		Instruction{Operation: Mul, X: 2, Y: 4},
		Instruction{Operation: Mul, X: 5, Y: 5},
		Instruction{Operation: Mul, X: 11, Y: 8},
		Instruction{Operation: Mul, X: 8, Y: 5},
	}

	if len(instructions) != len(expected) {
		fmt.Printf("got %v operations, expected %v\n", len(instructions), len(expected))
		t.Fail()
	}

	for i, op := range instructions {
		if op != expected[i] {
			fmt.Printf("got result %v, expected %v\n", op, expected[i])
			t.Fail()
		}
	}
}

func TestExecuteInstruction(t *testing.T) {
	i := Instruction{Operation: Mul, X: 2, Y: 4}
	result := ExecuteInstruction(i)
	expected := 8

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}

}

func TestDay3Part1(t *testing.T) {
	result := Part1(input1)
	expected := 161

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay3Part2(t *testing.T) {
	result := Part2(input2)
	expected := 48

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
