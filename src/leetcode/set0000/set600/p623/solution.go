package p623

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {

	var visit func(root *TreeNode, dep int)

	visit = func(root *TreeNode, dep int) {
		if root == nil {
			return
		}

		if dep == d-1 {
			left := &TreeNode{v, root.Left, nil}
			right := &TreeNode{v, nil, root.Right}

			root.Left = left
			root.Right = right
			return
		}

		visit(root.Left, dep+1)
		visit(root.Right, dep+1)
	}

	if d == 1 {
		newRoot := &TreeNode{v, root, nil}
		return newRoot
	}

	visit(root, 1)

	return root
}
