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
	arrLen := len(nums)
	// Handling border situations
	if arrLen == 0 {
		return 0
	}
	if arrLen == 1 {
		return nums[0]
	}

	dp0 := dps(nums[:arrLen-1], 0)
	dp1 := dps(nums[1:arrLen], 0)

	res1 := int((math.Max(float64(dp0[0]), float64(dp0[1]))))
	res2 := int((math.Max(float64(dp1[0]), float64(dp1[1]))))
	return int((math.Max(float64(res1), float64(res2))))
}

func main() {
	nums := []int{2, 3}
	fmt.Println(rob(nums))
}
