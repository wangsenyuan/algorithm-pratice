package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var travel func(root *TreeNode)

	var result []int
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		travel(root.Left)
		result = append(result, root.Val)
		travel(root.Right)
	}

	travel(root)

	return result
}
