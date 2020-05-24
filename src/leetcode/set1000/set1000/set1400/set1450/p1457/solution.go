package p1457

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {

	var dfs func(node *TreeNode)

	var res int

	cnt := make([]int, 10)

	dfs = func(node *TreeNode) {
		cnt[node.Val]++

		if node.Left == nil && node.Right == nil {
			var odd int

			for i := 1; i < 10; i++ {
				odd += cnt[i] & 1
			}

			if odd == 0 || odd == 1 {
				res++
			}
		} else {
			if node.Left != nil {
				dfs(node.Left)
			}

			if node.Right != nil {
				dfs(node.Right)
			}
		}

		cnt[node.Val]--
	}

	if root == nil {
		return 0
	}

	dfs(root)

	return res
}
