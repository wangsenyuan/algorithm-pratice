package p098

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var valid func(root *TreeNode) (int, int, bool)

	valid = func(root *TreeNode) (int, int, bool) {
		a, b, c := root.Val, root.Val, true
		var d int
		if root.Left != nil {
			a, d, c = valid(root.Left)
			if !c || d >= root.Val {
				return -1, -1, false
			}
		}

		if root.Right != nil {
			d, b, c = valid(root.Right)
			if !c || d <= root.Val {
				return -1, -1, false
			}
		}
		return a, b, c
	}

	_, _, c := valid(root)

	return c
}
