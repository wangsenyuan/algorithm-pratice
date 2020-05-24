package p1457

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {

	var dfs func(node *TreeNode, flag int)

	var res int

	dfs = func(node *TreeNode, flag int) {
		flag ^= (1 << node.Val)

		if node.Left == nil && node.Right == nil {
			if flag == 0 || flag == (flag&(-flag)) {
				res++
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, flag)
		}

		if node.Right != nil {
			dfs(node.Right, flag)
		}
	}

	if root == nil {
		return 0
	}

	dfs(root, 0)

	return res
}
