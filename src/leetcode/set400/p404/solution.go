package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum func(root *TreeNode, left bool) int

	sum = func(root *TreeNode, left bool) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			if left {
				return root.Val
			}
			return 0
		}

		return sum(root.Left, true) + sum(root.Right, false)
	}

	return sum(root, false)
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{6, nil, nil}}}
	fmt.Println(sumOfLeftLeaves(root))
}
