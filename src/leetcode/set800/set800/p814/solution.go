package p814

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {

	var prune func(root *TreeNode) (*TreeNode, int)

	prune = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}

		lt, ls := prune(root.Left)
		rt, rs := prune(root.Right)

		nr := &TreeNode{Val: root.Val}
		if ls != 0 {
			nr.Left = lt
		}
		if rs != 0 {
			nr.Right = rt
		}
		if ls+rs+root.Val > 0 {
			return nr, ls + rs + root.Val
		}
		return nil, 0
	}

	nr, _ := prune(root)

	return nr
}
