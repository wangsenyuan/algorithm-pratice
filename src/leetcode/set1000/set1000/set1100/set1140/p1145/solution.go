package p1145

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	cnt := make([]int, n+1)
	var xNode *TreeNode

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Val == x {
			xNode = node
		}

		cnt[node.Val] += dfs(node.Left)
		cnt[node.Val] += dfs(node.Right)
		cnt[node.Val]++

		return cnt[node.Val]
	}

	dfs(root)

	var a int
	if xNode.Left != nil {
		a = cnt[xNode.Left.Val]
	}
	var b int
	if xNode.Right != nil {
		b = cnt[xNode.Right.Val]
	}
	c := n - cnt[xNode.Val]

	if a*2 > n || b*2 > n || c*2 > n {
		return true
	}

	return false
}
