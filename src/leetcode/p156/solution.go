package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	left := upsideDownBinaryTree(root.Left)

	tmp := left
	for tmp.Right != nil {
		tmp = tmp.Right
	}

	tmp.Left, tmp.Right = root.Right, root
	root.Left, root.Right = nil, nil

	return left
}
