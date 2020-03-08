package p1373

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {

	if root == nil {
		return 0
	}

	const INF = math.MaxInt32
	const NINF = math.MinInt32

	var dfs func(node *TreeNode) []int

	var res int

	// return the [min, max] range of subtree rooted at node
	dfs = func(node *TreeNode) []int {
		if node.Left == nil && node.Right == nil {
			return []int{node.Val, node.Val, node.Val}
		}

		if node.Left != nil && node.Right != nil {
			a := dfs(node.Left)

			res = max(res, a[2])

			b := dfs(node.Right)

			res = max(res, b[2])

			if a[1] < a[0] {
				return []int{INF, NINF, NINF}
			}

			if b[1] < b[0] {
				return []int{INF, NINF, NINF}
			}

			if a[1] >= node.Val || b[0] <= node.Val {
				// not a valid BST
				return []int{INF, NINF, NINF}
			}

			//a[1] < node.Val && node.Val < b[0]
			return []int{a[0], b[1], a[2] + b[2] + node.Val}
		}

		if node.Left == nil {
			// only right child
			b := dfs(node.Right)
			res = max(res, b[2])

			if b[1] < b[0] {
				return []int{INF, NINF, NINF}
			}

			if b[0] <= node.Val {
				return []int{INF, NINF, NINF}
			}

			return []int{node.Val, b[1], node.Val + b[2]}
		}

		a := dfs(node.Left)
		res = max(res, a[2])

		if a[1] < a[0] {
			return []int{INF, NINF, NINF}
		}

		if a[1] >= node.Val {
			return []int{INF, NINF, NINF}
		}
		return []int{a[0], node.Val, node.Val + a[2]}
	}

	ans := dfs(root)

	if ans[0] > ans[1] {
		return res
	}

	return max(res, ans[2])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
