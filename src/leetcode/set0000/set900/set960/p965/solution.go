package p965

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	num := root.Val

	var loop func(node *TreeNode) bool

	loop = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != num {
			return false
		}

		return loop(node.Left) && loop(node.Right)
	}

	return loop(root)
}
