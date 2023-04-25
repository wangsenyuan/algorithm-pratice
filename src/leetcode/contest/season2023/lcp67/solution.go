package lcp67

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func expandBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := new(TreeNode)
	node.Val = root.Val
	if root.Left != nil {
		tmp := expandBinaryTree(root.Left)
		node.Left = new(TreeNode)
		node.Left.Val = -1
		node.Left.Left = tmp
	}

	if root.Right != nil {
		tmp := expandBinaryTree(root.Right)
		node.Right = new(TreeNode)
		node.Right.Val = -1
		node.Right.Right = tmp
	}

	return node
}
