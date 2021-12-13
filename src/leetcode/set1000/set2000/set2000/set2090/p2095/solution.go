package p2095

import "bytes"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	size := dfs1(root)
	flag := make([]int, size+1)

	var dfs2 func(node *TreeNode) int
	dfs2 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Val == startValue {
			flag[node.Val] |= 1
		}
		if node.Val == destValue {
			flag[node.Val] |= 2
		}
		flag[node.Val] |= dfs2(node.Left)
		flag[node.Val] |= dfs2(node.Right)
		return flag[node.Val]
	}

	dfs2(root)

	var buf bytes.Buffer
	var dfs3 func(node *TreeNode)
	var dfs4 func(node *TreeNode)

	dfs3 = func(node *TreeNode) {
		if node.Val == startValue {
			if flag[node.Val] == 3 {
				dfs4(node)
			}
			return
		}
		if node.Left != nil && flag[node.Left.Val] == 3 {
			dfs3(node.Left)
			return
		}
		if node.Right != nil && flag[node.Right.Val] == 3 {
			dfs3(node.Right)
			return
		}
		if node.Left != nil && flag[node.Left.Val] == 1 {
			dfs3(node.Left)
			buf.WriteByte('U')
			if node.Right != nil && flag[node.Right.Val] == 2 {
				buf.WriteByte('R')
				dfs4(node.Right)
			}
			return
		}
		// start has to be in right
		dfs3(node.Right)
		buf.WriteByte('U')
		if node.Left != nil && flag[node.Left.Val] == 2 {
			buf.WriteByte('L')
			dfs4(node.Left)
		}
	}

	dfs4 = func(node *TreeNode) {
		if node == nil || node.Val == destValue {
			return
		}

		if node.Left != nil && flag[node.Left.Val] == 2 {
			buf.WriteByte('L')
			dfs4(node.Left)
		} else {
			buf.WriteByte('R')
			dfs4(node.Right)
		}
	}

	dfs3(root)

	return buf.String()
}

func dfs1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + dfs1(root.Left) + dfs1(root.Right)
}
