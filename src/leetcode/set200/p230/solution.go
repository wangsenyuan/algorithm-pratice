package main

import "fmt"

func main() {
	root := &TreeNode{2, &TreeNode{1, nil, nil}, nil}
	fmt.Println(kthSmallest(root, 2))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ls := size(root.Left)
	if ls == k-1 {
		return root.Val
	}

	if ls > k-1 {
		return kthSmallest(root.Left, k)
	}

	return kthSmallest(root.Right, k-ls-1)
}

func size(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return size(root.Left) + size(root.Right) + 1
}
