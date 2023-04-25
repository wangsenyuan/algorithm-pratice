package lcp64

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closeLampInTree(root *TreeNode) int {
	type tuple struct {
		node             *TreeNode
		switch2, switch3 bool
	}
	dp := map[tuple]int{} // 记忆化搜索
	var f func(*TreeNode, bool, bool) int
	f = func(node *TreeNode, switch2, switch3 bool) int {
		if node == nil {
			return 0
		}
		p := tuple{node, switch2, switch3}
		if res, ok := dp[p]; ok {
			return res
		}
		if node.Val == 1 == (switch2 == switch3) { // 当前节点为开灯
			res1 := f(node.Left, switch2, false) + f(node.Right, switch2, false) + 1
			res2 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 1
			res3 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 1
			r123 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 3
			dp[p] = min(res1, res2, res3, r123)
		} else { // 当前节点为关灯
			res0 := f(node.Left, switch2, false) + f(node.Right, switch2, false)
			res12 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 2
			res13 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 2
			res23 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 2
			dp[p] = min(res0, res12, res13, res23)
		}
		return dp[p]
	}
	return f(root, false, false)
}

const INF = 1 << 30

func min(a, b, c, d int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	if d < a {
		a = d
	}
	return a
}
