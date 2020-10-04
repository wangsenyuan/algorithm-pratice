package p1528

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) []int

	var ans int

	dfs = func(node *TreeNode) []int {
		res := make([]int, distance+1)
		if node.Left == nil && node.Right == nil {
			// a leaf
			res[0] = 1
			return res
		}

		if node.Left == nil || node.Right == nil {
			var tmp []int
			if node.Left == nil {
				tmp = dfs(node.Right)
			} else {
				tmp = dfs(node.Left)
			}
			update(res, tmp, distance)
			return res
		}

		// have both left & right
		a := dfs(node.Left)
		update(a, a, distance)
		b := dfs(node.Right)
		update(b, b, distance)

		for x := 1; x <= distance; x++ {
			for y := 1; x+y <= distance; y++ {
				ans += a[x] * b[y]
			}
		}

		for x := 1; x <= distance; x++ {
			res[x] += a[x]
			res[x] += b[x]
		}
		return res
	}
	dfs(root)

	return ans
}

func update(a, b []int, dist int) {
	for i := dist - 1; i >= 0; i-- {
		a[i+1] = b[i]
	}
}
