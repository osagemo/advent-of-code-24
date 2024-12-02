package main

import (
	"fmt"
	"testing"
)

const input1 = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
9 1 2 3 4
1 9 2 3 4
1 2 9 4 5
1 2 3 9 5
1 2 3 4 9`

func TestDay2Part1(t *testing.T) {
	result := Part1(input1)
	expected := 2

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay2Part2(t *testing.T) {
	result := Part2(input1)
	expected := 9

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
