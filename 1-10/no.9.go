package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}
	_x := x
	r := 0
	for _x > 0 {
		r = r*10 + _x%10
		_x /= 10
	}

	return x == r
}

func main() {
	fmt.Println(isPalindrome(-0))
}
