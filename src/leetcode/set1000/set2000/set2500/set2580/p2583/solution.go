package p2583

import "sort"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var sums []int64

	var dfs func(node *TreeNode, d int)

	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if len(sums) == d {
			sums = append(sums, 0)
		}
		sums[d] += int64(node.Val)

		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}

	dfs(root, 0)

	if len(sums) < k {
		return -1
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	return sums[k-1]
}
