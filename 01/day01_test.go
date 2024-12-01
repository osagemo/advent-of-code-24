package main

import (
	"fmt"
	"testing"
)

const input1 = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1Part1(t *testing.T) {
	result := Part1(input1)
	expected := 11

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	result := Part2(input1)
	expected := 31

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
