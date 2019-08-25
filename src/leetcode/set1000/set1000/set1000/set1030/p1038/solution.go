package p1038

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	var loop func(node *TreeNode, carry int) (*TreeNode, int)

	loop = func(node *TreeNode, carry int) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}

		right, tmp := loop(node.Right, carry)
		newNode := new(TreeNode)
		newNode.Right = right
		newNode.Val = carry + node.Val + tmp

		left, tmp2 := loop(node.Left, newNode.Val)
		newNode.Left = left
		return newNode, node.Val + tmp + tmp2
	}

	res, _ := loop(root, 0)
	return res
}
