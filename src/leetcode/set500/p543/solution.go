package main

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	var visit func(root *TreeNode) (int, int)

	visit = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		lh, ld := visit(root.Left)
		rh, rd := visit(root.Right)

		h := max(lh, rh) + 1
		d := max(lh+rh, max(ld, rd))
		return h, d
	}

	_, d := visit(root)

	return d
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
