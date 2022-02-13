package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func countBinarySubstrings(s string) int {
	res, last, ptr, n := 0, 0, 0, len(s)

	for ptr < n {
		count := 0
		ch := s[ptr]

		for ptr < n && ch == s[ptr] {
			count++
			ptr++
		}

		res += min(count, last)
		last = count
	}

	return res
}

func main() {
	s1 := "111001"
	fmt.Println(countBinarySubstrings(s))
}
