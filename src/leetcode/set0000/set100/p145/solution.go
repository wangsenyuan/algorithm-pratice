package p145

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int

	a := root

	for a != nil {
		if a.Right != nil {
			b := a.Right
			for b.Left != nil && b.Left != a {
				b = b.Left
			}

			if b.Left == a {
				b.Left = nil
				a = a.Left
			} else {
				res = append(res, a.Val)
				b.Left = a
				a = a.Right
			}
		} else {
			res = append(res, a.Val)
			a = a.Left
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
