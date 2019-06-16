package p103

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var visit func(nodes []*TreeNode, l2r bool)
	result := make([][]int, 0, 10)

	visit = func(nodes []*TreeNode, l2r bool) {
		if len(nodes) == 0 {
			return
		}
		n := len(nodes)

		vals := make([]int, 0, 10)
		for i := 0; i < n; i++ {
			vals = append(vals, nodes[i].Val)
		}

		result = append(result, vals)

		next := make([]*TreeNode, 0, 10)
		for i := n - 1; i >= 0; i-- {
			node := nodes[i]
			if l2r {
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			} else {
				if node.Right != nil {
					next = append(next, node.Right)
				}
				if node.Left != nil {
					next = append(next, node.Left)
				}
			}
		}
		visit(next, !l2r)
	}

	visit([]*TreeNode{root}, false)
	return result
}
