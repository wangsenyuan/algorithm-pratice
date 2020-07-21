package p101

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var isMirror func(left, right *TreeNode) bool

	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return root == nil || isMirror(root.Left, root.Right)
}
