package main

import "fmt"

const MAX = 2147483648 - 1
const MIN = -2147483648

func reverse(x int) int {
	var arr [10] int
	var minus = x < 0

	if minus {
		x *= -1
	}

	res := 0
	_len := 0
	for ;x > 0; {
		arr[_len] = x % 10
		x = x / 10
		_len ++
	}
	for i := 0; i < _len; i++ {
		if i > 0 {
			res *= 10
		}

		res += arr[i]
	}

	if minus {
		res *= -1
	}

	if minus && res < MIN {
		return 0
	} else if res > MAX {
		return 0
	}

	return res
}

func main() {
	fmt.Println(reverse(120))
}
