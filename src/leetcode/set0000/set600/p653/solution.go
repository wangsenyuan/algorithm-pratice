package p653

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {

	var find func(root *TreeNode, x int) *TreeNode

	find = func(root *TreeNode, x int) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == x {
			return root
		}

		if x < root.Val {
			return find(root.Left, x)
		}

		return find(root.Right, x)
	}

	var visit func(root *TreeNode) bool

	visit = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		x := k - node.Val

		candidate := find(root, x)

		if candidate != nil && candidate != node {
			return true
		}

		return visit(node.Left) || visit(node.Right)
	}

	return visit(root)
}
