package p113

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(dfs(node.Left), dfs(node.Right)) + 1
	}

	h := dfs(root)

	path := make([]int, h)

	var res [][]int

	var dfs2 func(node *TreeNode, cur int, i int)

	dfs2 = func(node *TreeNode, cur int, i int) {
		if node == nil {
			return
		}
		path[i] = node.Val
		cur += node.Val
		if node.Left == nil && node.Right == nil {
			if cur == sum {
				tmp := make([]int, i+1)
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		dfs2(node.Left, cur, i+1)
		dfs2(node.Right, cur, i+1)
	}

	dfs2(root, 0, 0)

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
