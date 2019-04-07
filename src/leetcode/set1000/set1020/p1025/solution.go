package p1025

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	const mod = 1e9 + 7
	var res int64
	var loop func(node *TreeNode, cur int64)

	loop = func(node *TreeNode, cur int64) {
		cur = (cur<<1 | int64(node.Val)) % mod
		if node.Left == nil && node.Right == nil {
			res += cur
			if res >= mod {
				res -= mod
			}
			return
		}
		if node.Left != nil {
			loop(node.Left, cur)
		}

		if node.Right != nil {
			loop(node.Right, cur)
		}
	}

	loop(root, 0)

	return int(res)
}
