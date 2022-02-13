package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var numsMap = make(map[int]int)
	var left, right int

	for i, num := range nums {
		numsMap[num] = i
	}

	for i, num := range nums {
		k := target - num

		if j, ok := numsMap[k]; ok && j != i {
			left = i
			right = j
			break
		}
	}

	return []int{left, right}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}
