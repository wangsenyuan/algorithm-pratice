package p530

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	findLeftMax := func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		for root.Right != nil {
			root = root.Right
		}
		return root.Val
	}

	findRightMin := func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		for root.Left != nil {
			root = root.Left
		}
		return root.Val
	}

	var visit func(root *TreeNode)
	var res = math.MaxInt32

	visit = func(root *TreeNode) {
		if root == nil {
			return
		}

		leftMax := findLeftMax(root.Left)
		rightMin := findRightMin(root.Right)

		if leftMax != -1 && root.Val-leftMax < res {
			res = root.Val - leftMax
		}

		if rightMin != -1 && rightMin-root.Val < res {
			res = rightMin - root.Val
		}
		visit(root.Left)
		visit(root.Right)
	}

	visit(root)

	return res
}
