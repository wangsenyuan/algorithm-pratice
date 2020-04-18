package e

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimalExecTime(root *TreeNode) float64 {

	var dfs func(node *TreeNode) (int, float64)

	dfs = func(node *TreeNode) (int, float64) {
		if node == nil {
			return 0, 0
		}

		a, b := dfs(node.Left)
		c, d := dfs(node.Right)

		e := math.Max(b, d)
		e = math.Max(e, float64(a+c)/2)
		e += float64(node.Val)

		return a + c + node.Val, e
	}

	_, res := dfs(root)

	return res
}
