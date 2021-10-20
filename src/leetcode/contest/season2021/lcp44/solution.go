package lcp44

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {
	cnt := make(map[int]int)

	var loop func(node *TreeNode)
	loop = func(node *TreeNode) {
		if node != nil {
			cnt[node.Val]++
			loop(node.Left)
			loop(node.Right)
		}
	}

	loop(root)

	return len(cnt)
}
