package p889_

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) != len(post) {
		return nil
	}
	if len(pre) == 0 {
		return nil
	}
	n := len(pre)
	root := new(TreeNode)
	root.Val = pre[0]
	if n == 1 {
		return root
	}

	for i := 0; i < n-1; i++ {
		if post[i] == pre[1] {
			root.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
			root.Right = constructFromPrePost(pre[i+2:], post[i+1:n-1])
			break
		}
	}
	return root
}

func preOrderTraversal(root *TreeNode) []int {
	var travel func(root *TreeNode)

	var result []int
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		travel(root.Left)
		travel(root.Right)
	}

	travel(root)

	return result
}

func postOrderTraversal(root *TreeNode) []int {
	var travel func(root *TreeNode)

	var result []int
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		travel(root.Right)
		result = append(result, root.Val)
	}

	travel(root)

	return result
}
