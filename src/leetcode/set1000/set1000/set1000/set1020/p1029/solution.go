package p1029

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var best int

	check := func(a, b, c int, min, max *int) {
		if abs(a-c) > best {
			best = abs(a - c)
		}
		if abs(b-c) > best {
			best = abs(b - c)
		}
		if a < *min {
			*min = a
		}
		if b > *max {
			*max = b
		}
	}

	var loop func(node *TreeNode) (min int, max int)

	loop = func(node *TreeNode) (min int, max int) {
		min = node.Val
		max = node.Val
		if node.Left == nil && node.Right == nil {
			min = node.Val
			max = node.Val
			return
		}

		if node.Left != nil {
			a, b := loop(node.Left)
			check(a, b, node.Val, &min, &max)
		}

		if node.Right != nil {
			a, b := loop(node.Right)
			check(a, b, node.Val, &min, &max)
		}

		return
	}

	loop(root)

	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
