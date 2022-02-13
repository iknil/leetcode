package main

import "fmt"

func getInitArray() *[128]int {
	var res [128]int

	for i, _ := range res {
		res[i] = -1
	}

	return &res
}

func lengthOfLongestSubstring(s string) int {
	_len := 0
	_s := 0
	_l := 0
	wordMap := getInitArray()

	for i, val := range s {
		if j := wordMap[val]; j >= _s {
			_l = i - _s
			if _l > _len {
				_len = _l
			}
			_s = j + 1
		}

		wordMap[val] = i
	}

	_l = len(s) - _s
	if _l > _len {
		_len = _l
	}

	return _len
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}
