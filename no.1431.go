package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0

	for _, num := range candies {
		if num > max {
			max = num
		}
	}

	for i, num := range candies {
		if num + extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

func main() {
	candies := []int{12,1,12}
	extraCandies := 10

	fmt.Println(kidsWithCandies(candies, extraCandies))
}