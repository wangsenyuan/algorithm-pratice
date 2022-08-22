package p2385

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	height := make(map[int]int)
	max_height := make(map[int]int)

	var dfs1 func(node *TreeNode, h int) int

	dfs1 = func(node *TreeNode, h int) int {
		if node == nil {
			return h - 1
		}
		height[node.Val] = h
		xh := max(dfs1(node.Left, h+1), dfs1(node.Right, h+1))
		max_height[node.Val] = xh
		return xh
	}

	dfs1(root, 0)

	where := make(map[int]int)

	var dfs2 func(node *TreeNode) bool

	dfs2 = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if node.Val == start {
			return true
		}
		a := dfs2(node.Left)
		if a {
			where[node.Val] = 1
			return true
		}
		b := dfs2(node.Right)
		if b {
			where[node.Val] = 2
			return true
		}
		return false
	}

	dfs2(root)

	var best int

	var dfs3 func(node *TreeNode)

	dfs3 = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val == start {
			best = max(best, max_height[node.Val]-height[node.Val])
			return
		}

		x := where[node.Val]
		h := height[start] - height[node.Val]

		if x == 1 {
			// go left
			if node.Right != nil {
				h += max_height[node.Right.Val] - height[node.Val]
			}
			best = max(best, h)
			dfs3(node.Left)
		} else {
			// go right
			if node.Left != nil {
				h += max_height[node.Left.Val] - height[node.Val]
			}
			best = max(best, h)

			dfs3(node.Right)
		}
	}

	dfs3(root)

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
