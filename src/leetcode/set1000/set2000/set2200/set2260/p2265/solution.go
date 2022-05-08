package p2265

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {

	var res int

	var dfs func(root *TreeNode) (int, int)

	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		sum := root.Val
		cnt := 1

		a, b := dfs(root.Left)
		sum += a
		cnt += b

		a, b = dfs(root.Right)

		sum += a
		cnt += b

		if root.Val == sum/cnt {
			res++
		}
		return sum, cnt
	}

	dfs(root)

	return res
}
