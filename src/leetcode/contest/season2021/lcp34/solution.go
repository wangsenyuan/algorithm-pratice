package lcp34

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {

	var loop func(node *TreeNode) []int

	loop = func(node *TreeNode) []int {
		res := make([]int, k+1)

		if node == nil {
			return res
		}

		a := loop(node.Left)
		b := loop(node.Right)

		for i := 0; i <= k; i++ {
			for j := 0; j <= k; j++ {
				res[0] = max(res[0], a[i]+b[j])
				if i+j+1 <= k {
					res[i+j+1] = max(res[i+j+1], a[i]+b[j]+node.Val)
				}
			}
		}

		return res
	}

	res := loop(root)
	best := res[0]
	for i := 0; i <= k; i++ {
		best = max(best, res[i])
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
