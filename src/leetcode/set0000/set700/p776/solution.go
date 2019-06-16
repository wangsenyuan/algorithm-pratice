package p776

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, V int) []*TreeNode {
	ans := split(root, V)
	if ans[0] == nil {
		return ans[1:]
	}
	return ans
}

func split(root *TreeNode, V int) []*TreeNode {

	var find func(root *TreeNode) bool

	find = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if root.Val <= V {
			return true
		}
		if V < root.Val {
			return find(root.Left)
		}
		return find(root.Right)
	}

	var visit func(root *TreeNode, parentGt bool) *TreeNode

	ans := make([]*TreeNode, 0, 2)

	visit = func(root *TreeNode, parentGt bool) *TreeNode {
		if root == nil {
			return root
		}
		if root.Val <= V && parentGt {
			ans = append(ans, root)
			tmp := split(root.Right, V)
			root.Right = nil
			if len(tmp) > 0 {
				root.Right = tmp[0]
			}
			if len(tmp) > 1 {
				return tmp[1]
			}
			return nil
		}

		root.Left = visit(root.Left, root.Val > V)
		if V > root.Val {
			root.Right = visit(root.Right, false)
		}
		return root
	}

	if find(root) {
		tmp := visit(root, false)
		if tmp != nil {
			ans = append(ans, tmp)
		}
	} else {
		ans = []*TreeNode{nil, root}
	}

	return ans
}
