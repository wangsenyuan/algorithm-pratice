package p1302

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max_d := depth(root)

	var dfs func(node *TreeNode, d int)
	var sum int

	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}

		if d == max_d {
			sum += node.Val
			return
		}

		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}

	dfs(root, 1)

	return sum
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
