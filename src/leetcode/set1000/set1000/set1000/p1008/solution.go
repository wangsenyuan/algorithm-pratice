package p1008

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]

	var i = 1
	for i < len(preorder) && preorder[i] < preorder[0] {
		i++
	}
	root.Left = bstFromPreorder(preorder[1:i])
	root.Right = bstFromPreorder(preorder[i:])
	return root
}
