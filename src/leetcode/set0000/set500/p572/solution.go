package p572

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}

	if s == nil {
		return false
	}

	if isSameTree(s, t) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
