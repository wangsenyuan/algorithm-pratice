package p1343

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {

	var sum int64

	var res int64

	var dfs func(node *TreeNode) int64
	dfs = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}
		x := int64(node.Val) + dfs(node.Left) + dfs(node.Right)
		y := sum - x

		res = max(res, x*y)

		return x
	}

	sum = dfs(root)

	res = 0

	dfs(root)

	return int(res % 1000000007)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
