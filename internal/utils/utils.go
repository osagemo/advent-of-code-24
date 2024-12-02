package utils

import "strconv"

func MustParseInts(strings []string) []int {
	ints := []int{}
	for _, s := range strings {
		ints = append(ints, MustParseInt(s))
	}
	return ints
}

func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
