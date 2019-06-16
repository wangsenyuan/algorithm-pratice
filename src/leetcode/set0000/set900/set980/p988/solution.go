package p988

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const s = "abcdefghijklmnopqrstuvwxyz"

func smallestFromLeaf(root *TreeNode) string {

	var dfs func(node *TreeNode) string

	dfs = func(node *TreeNode) string {
		x := string(s[node.Val])

		if node.Left == nil && node.Right == nil {
			return x
		}

		if node.Right == nil {
			return dfs(node.Left) + x
		}

		if node.Left == nil {
			return dfs(node.Right) + x
		}

		left := dfs(node.Left) + x
		right := dfs(node.Right) + x
		if left < right {
			return left
		}
		return right
	}

	if root == nil {
		return ""
	}
	return dfs(root)
}
