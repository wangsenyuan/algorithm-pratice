package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {

	var change func(root *TreeNode, gt int) (*TreeNode, int)

	change = func(root *TreeNode, gt int) (*TreeNode, int) {
		if root == nil {
			return nil, gt
		}

		right, rgt := change(root.Right, gt)

		left, lgt := change(root.Left, rgt+root.Val)

		return &TreeNode{rgt + root.Val, left, right}, lgt
	}

	newRoot, _ := change(root, 0)

	return newRoot
}
