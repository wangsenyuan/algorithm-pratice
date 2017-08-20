package main

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftPos := make(map[int]int)
	ans := 0
	var visit func(node *TreeNode, h int, i int)

	visit = func(node *TreeNode, h int, i int) {
		if node == nil {
			return
		}

		if pos, found := leftPos[h]; found {
			if i-pos+1 > ans {
				ans = i - pos + 1
			}
		} else {
			leftPos[h] = i
		}

		visit(node.Left, h+1, 2*i)
		visit(node.Right, h+1, 2*i+1)
	}

	leftPos[0] = 1
	visit(root, 0, 1)

	return ans
}
