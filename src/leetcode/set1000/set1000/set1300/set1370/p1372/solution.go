package p1372

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {

	mem := make(map[*TreeNode][]int)

	var dfs2 func(node *TreeNode, dir int) int

	dfs2 = func(node *TreeNode, dir int) int {
		if node == nil {
			return 0
		}

		if mem[node] != nil && mem[node][dir] > 0 {
			return mem[node][dir]
		}
		if mem[node] == nil {
			mem[node] = make([]int, 2)
		}

		if dir == 0 {
			mem[node][dir] = 1 + dfs2(node.Left, 1)
		} else {
			mem[node][dir] = 1 + dfs2(node.Right, 0)
		}

		return mem[node][dir]
	}

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		ans := dfs2(node.Left, 1) + 1
		ans = max(ans, dfs2(node.Right, 0)+1)
		ans = max(ans, dfs(node.Left))
		ans = max(ans, dfs(node.Right))

		return ans
	}

	if root == nil {
		return 0
	}

	return dfs(root) - 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
