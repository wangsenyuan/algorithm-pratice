package p993

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	level := make([]int, 101)
	parent := make([]int, 101)

	var dfs func(node *TreeNode, p int, l int)

	dfs = func(node *TreeNode, p int, l int) {
		if node == nil {
			return
		}
		level[node.Val] = l
		parent[node.Val] = p
		dfs(node.Left, node.Val, l+1)
		dfs(node.Right, node.Val, l+1)
	}
	dfs(root, 0, 0)

	return level[x] == level[y] && parent[x] != parent[y]
}
