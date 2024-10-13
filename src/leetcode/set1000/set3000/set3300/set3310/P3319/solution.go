package p3319

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {

	var arr []int

	var dfs func(node *TreeNode) (bool, int)

	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		ok1, l := dfs(node.Left)
		ok2, r := dfs(node.Right)
		sz := l + r + 1

		var ok bool

		if l == r && ok1 && ok2 {
			ok = true
			arr = append(arr, sz)
		}

		return ok, sz
	}

	dfs(root)

	sort.Ints(arr)

	n := len(arr)
	if n < k {
		return -1
	}

	return arr[n-k]
}
