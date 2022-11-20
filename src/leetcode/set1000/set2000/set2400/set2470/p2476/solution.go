package p2476

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {

	var vals []int

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}
		vals = append(vals, node.Val)
		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)

	res := make([][]int, len(queries))

	for i, cur := range queries {
		res[i] = []int{-1, -1}

		x := search(len(vals), func(j int) bool {
			return vals[j] > cur
		})

		if x > 0 {
			res[i][0] = vals[x-1]
		}

		y := search(len(vals), func(j int) bool {
			return vals[j] >= cur
		})
		if y < len(vals) {
			res[i][1] = vals[y]
		}
	}

	return res
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
