package p106

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := postorder[len(postorder)-1]

	i := len(inorder) - 1
	for i >= 0 && inorder[i] != root {
		i--
	}

	return &TreeNode{root, buildTree(inorder[:i], postorder[:i]), buildTree(inorder[i+1:], postorder[i:len(postorder)-1])}
}
