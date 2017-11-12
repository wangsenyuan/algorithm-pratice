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

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	mid := n / 2

	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func main() {
	nums := []int{3, 5, 8}
	root := sortedArrayToBST(nums)
	output(root)
}

func output(root *TreeNode) {
	if root == nil {
		fmt.Print("_")
		return
	}

	fmt.Printf("%d{", root.Val)
	output(root.Left)
	output(root.Right)
	fmt.Print("}")
}
