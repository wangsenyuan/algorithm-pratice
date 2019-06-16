package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := preorder[0]

	i := 0
	for i < len(inorder) && inorder[i] != root {
		i++
	}

	return &TreeNode{root, buildTree(preorder[1:i+1], inorder[:i]), buildTree(preorder[i+1:], inorder[i+1:])}
}
