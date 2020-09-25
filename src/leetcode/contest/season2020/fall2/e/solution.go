package e

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func navigation(root *TreeNode) int {
	if root == nil {
		return 0
	}

	c1, n1 := dfs(root.Left)
	c2, n2 := dfs(root.Right)

	res := c1 + c2

	if n1 && c2 == 0 || n2 && c1 == 0 {
		res++
	}
	return res
}

func dfs(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	if node.Left == nil && node.Right == nil {
		return 0, true
	}
	c1, n1 := dfs(node.Left)
	c2, n2 := dfs(node.Right)

	c := c1 + c2

	n := (n1 && c2 == 0) || (n2 && c1 == 0)

	if node.Left != nil && node.Right != nil && c1 == 0 && c2 == 0 {
		c++
	}

	if node.Left != nil && node.Right != nil && (c1 == 0 || c2 == 0) {
		n = true
	}

	return c, n
}
