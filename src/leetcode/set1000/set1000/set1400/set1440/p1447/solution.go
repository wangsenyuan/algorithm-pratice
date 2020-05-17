package p1447

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	var dfs func(node *TreeNode, x int)

	var res int

	dfs = func(node *TreeNode, x int) {
		if node == nil {
			return
		}
		if node.Val >= x {
			res++
		}
		dfs(node.Left, max(x, node.Val))
		dfs(node.Right, max(x, node.Val))
	}

	dfs(root, math.MinInt32)

	return res

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
