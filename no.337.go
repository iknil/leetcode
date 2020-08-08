package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	res := [2]int{0, 0}

	res[0] = int(math.Max(float64(left[0]), float64(left[1])) + math.Max(float64(right[0]), float64(right[1])))
	res[1] = root.Val + left[0] + right[0]

	return res
}

func rob(root *TreeNode) int {
	res := dfs(root)
	return int(math.Max(float64(res[0]), float64(res[1])))
}

func main() {
	var root *TreeNode = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	fmt.Println(rob(root))
}
