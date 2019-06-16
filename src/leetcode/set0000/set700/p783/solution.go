package p783

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	const INF = 1 << 18

	// return diff, min, max
	var visit func(root *TreeNode) (int, int, int)

	visit = func(root *TreeNode) (int, int, int) {
		if root == nil {
			return INF, INF, -INF
		}

		diff, min, max := INF, root.Val, root.Val
		if root.Left != nil {
			a, b, c := visit(root.Left)
			d := root.Val - c
			// left children all less than root
			diff = minValue(a, d)
			min = b
		}

		if root.Right != nil {
			a, b, c := visit(root.Right)
			d := b - root.Val
			diff = minValue(diff, minValue(a, d))
			max = maxValue(max, c)
		}
		return diff, min, max
	}
	res, _, _ := visit(root)

	return res
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
