package p998

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		res := new(TreeNode)
		res.Val = val
		return res
	}
	if root.Val < val {
		res := new(TreeNode)
		res.Val = val
		res.Left = root
		return res
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
