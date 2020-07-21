package p124

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32

	var process func(*TreeNode) int

	process = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, process(root.Left))
		right := max(0, process(root.Right))
		result = max(result, left+right+root.Val)
		return max(left, right) + root.Val
	}

	process(root)

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
