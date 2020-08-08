package main

import (
	"fmt"
	"math"
)

func dps(nums []int, index int) [2]int {
	if index >= len(nums) {
		return [2]int{0, 0}
	}

	dp := [2]int{0, 0}
	next := dps(nums, index+1)
	dp[0] = int(math.Max(float64(next[0]), float64(next[1])))
	dp[1] = next[0] + nums[index]

	return dp
}

func rob(nums []int) int {
	res := dps(nums, 0)
	return int((math.Max(float64(res[0]), float64(res[1]))))
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
