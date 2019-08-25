package p1110

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	set := make(map[int]bool)

	for _, num := range to_delete {
		set[num] = true
	}
	var res []*TreeNode

	var loop func(parent *TreeNode, node *TreeNode, flag bool)

	loop = func(parent, node *TreeNode, flag bool) {
		if node == nil {
			return
		}
		if set[node.Val] {
			if parent != nil {
				if parent.Left == node {
					parent.Left = nil
				}
				if parent.Right == node {
					parent.Right = nil
				}
			}
			loop(node, node.Left, true)
			loop(node, node.Right, true)
			return
		}
		if flag {
			res = append(res, node)
			flag = false
		}
		loop(node, node.Left, flag)
		loop(node, node.Right, flag)
	}

	loop(nil, root, true)

	return res
}
