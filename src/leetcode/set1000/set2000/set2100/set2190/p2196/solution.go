package p2196

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode)
	parent := make(map[int]int)
	child := -1
	for _, des := range descriptions {
		p, c, l := des[0], des[1], des[2]
		if nodes[p] == nil {
			nodes[p] = new(TreeNode)
			nodes[p].Val = p
		}
		parent[c] = p
		child = c
		if nodes[c] == nil {
			nodes[c] = new(TreeNode)
			nodes[c].Val = c
		}
		if l == 1 {
			nodes[p].Left = nodes[c]
		} else {
			nodes[p].Right = nodes[c]
		}
	}

	for {
		if p, ok := parent[child]; !ok {
			return nodes[child]
		} else {
			child = p
		}
	}
}
