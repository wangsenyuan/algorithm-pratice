package p1031

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {

	var loop func(s string, d int) *TreeNode

	loop = func(s string, d int) *TreeNode {
		// find right child of root, that has d + 1 dashs
		if len(s) <= d {
			// invalid?
			return nil
		}
		s = s[d:]
		var num int
		var i int
		for i < len(s) && s[i] != '-' {
			num = num*10 + int(s[i]-'0')
			i++
		}
		node := new(TreeNode)
		node.Val = num

		//following left child if any
		if i == len(s) {
			return node
		}

		j := i + d + 1
		var u int
		for j < len(s) {
			if s[j] == '-' {
				u++
			} else {
				if u == d+1 {
					// found
					break
				}
				u = 0
			}

			j++
		}
		node.Left = loop(s[i:j-u], d+1)
		node.Right = loop(s[j-u:], d+1)
		return node
	}

	return loop(S, 0)
}
