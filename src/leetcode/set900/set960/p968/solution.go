package p968

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAX = math.MaxInt32 >> 2

func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var loop func(node *TreeNode) []int

	loop = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0, MAX}
		}

		L := loop(node.Left)
		R := loop(node.Right)

		a := min(L[1], L[2])
		b := min(R[1], R[2])
		d0 := L[1] + R[1]
		d1 := min(R[2]+a, L[2]+b)
		d2 := 1 + min(L[0], a) + min(R[0], b)
		return []int{d0, d1, d2}
	}

	res := loop(root)
	return min(res[1], res[2])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
