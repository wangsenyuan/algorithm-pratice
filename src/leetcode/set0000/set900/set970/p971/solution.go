package p971

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	n := len(voyage)

	ord := make([]int, n+1)
	for i := 0; i < n; i++ {
		ord[voyage[i]] = i
	}

	var dfs func(node *TreeNode) bool
	open := make([]int, n+1)
	close := make([]int, n+1)
	res := make([]int, 0, n)
	dfs = func(node *TreeNode) bool {
		if node != nil {
			open[node.Val] = ord[node.Val]
			close[node.Val] = ord[node.Val]
			if !dfs(node.Left) {
				return false
			}
			if !dfs(node.Right) {
				return false
			}
			if node.Left != nil && ord[node.Left.Val] < ord[node.Val] {
				return false
			}
			if node.Right != nil && ord[node.Right.Val] < ord[node.Val] {
				return false
			}

			if node.Left != nil && node.Right != nil {
				if open[node.Right.Val] > close[node.Left.Val] {
					//no need to swap
				} else if open[node.Left.Val] > close[node.Right.Val] {
					//need to swap
					res = append(res, node.Val)
				} else {
					//overlap
					return false
				}
			}
			if node.Left != nil {
				open[node.Val] = min(open[node.Val], open[node.Left.Val])
				close[node.Val] = max(close[node.Val], close[node.Left.Val])
			}
			if node.Right != nil {
				open[node.Val] = min(open[node.Val], open[node.Right.Val])
				close[node.Val] = max(close[node.Val], close[node.Right.Val])
			}
		}

		return true
	}

	if dfs(root) {
		return res
	}

	return []int{-1}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
