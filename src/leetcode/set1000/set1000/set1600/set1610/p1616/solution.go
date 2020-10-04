package p1616

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	h := depth(root)
	vs := make([]int, h)
	for i := 0; i < h; i++ {
		if i&1 == 0 {
			vs[i] = math.MinInt32
		} else {
			vs[i] = math.MaxInt32
		}
	}

	var dfs func(node *TreeNode, d int) bool
	dfs = func(node *TreeNode, d int) bool {
		if node == nil {
			return true
		}

		if d&1 == 0 {
			if node.Val&1 == 0 || node.Val <= vs[d] {
				return false
			}
		} else {
			if node.Val&1 == 1 || node.Val >= vs[d] {
				return false
			}
		}
		vs[d] = node.Val

		return dfs(node.Left, d+1) && dfs(node.Right, d+1)
	}

	return dfs(root, 0)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
