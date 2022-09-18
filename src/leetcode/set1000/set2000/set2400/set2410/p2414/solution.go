package p2414

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	var h int

	node := root

	for node != nil {
		h++
		node = node.Left
	}

	sz := 1 << h
	arr := make([]int, sz-1)

	var dfs func(node *TreeNode, id int)

	dfs = func(node *TreeNode, id int) {
		if node == nil {
			return
		}
		arr[id] = node.Val
		dfs(node.Left, 2*id+1)
		dfs(node.Right, 2*id+2)
	}

	dfs(root, 0)

	for i := 1; i < h; i += 2 {
		l := (1 << i) - 1
		r := (1 << (i + 1)) - 2
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	var dfs2 func(id int) *TreeNode

	dfs2 = func(id int) *TreeNode {
		if id >= sz-1 {
			return nil
		}
		node := &TreeNode{Val: arr[id]}
		node.Left = dfs2(2*id + 1)
		node.Right = dfs2(2*id + 2)
		return node
	}

	return dfs2(0)
}
