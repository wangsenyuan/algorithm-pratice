package recoverfrompreoder

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	n := len(S)
	var dfs func(level int, pos int) (*TreeNode, int)

	dfs = func(level int, pos int) (*TreeNode, int) {
		var depth int

		for pos+depth < n && S[pos+depth] == '-' {
			depth++
		}

		if depth != level {
			return nil, 0
		}

		var num int
		var i int = pos + depth
		for i < n && (S[i] >= '0' && S[i] <= '9') {
			num = num*10 + int(S[i]-'0')
			i++
		}
		node := new(TreeNode)
		node.Val = num

		tmpLeft, j := dfs(level+1, i)
		if tmpLeft != nil {
			node.Left = tmpLeft
			tmpRight, jj := dfs(level+1, j)
			if tmpRight == nil {
				return node, j
			}
			node.Right = tmpRight
			return node, jj
		}
		return node, i
	}

	node, _ := dfs(0, 0)

	return node
}
