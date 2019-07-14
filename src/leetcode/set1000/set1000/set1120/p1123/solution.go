package p1123

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	dd := make(map[*TreeNode]int)

	var dfs1 func(node *TreeNode) int

	dfs1 = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		a := dfs1(node.Left)
		b := dfs1(node.Right)
		dd[node] = max(a, b) + 1
		return dd[node]
	}

	dfs1(root)

	var dfs2 func(node *TreeNode) *TreeNode

	dfs2 = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Left != nil && node.Right != nil {
			if dd[node.Left] == dd[node.Right] {
				return node
			}
			if dd[node.Left] > dd[node.Right] {
				return dfs2(node.Left)
			}
			return dfs2(node.Right)
		}
		if node.Left != nil {
			return dfs2(node.Left)
		}
		if node.Right != nil {
			return dfs2(node.Right)
		}
		return node
	}

	return dfs2(root)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
