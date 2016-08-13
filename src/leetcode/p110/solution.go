package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := height(root.Left), height(root.Right)
	if l-r > 1 || r-l > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(height(root.Left), height(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
