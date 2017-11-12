package main

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxAt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxAt] {
			maxAt = i
		}
	}

	left := constructMaximumBinaryTree(nums[:maxAt])
	right := constructMaximumBinaryTree(nums[maxAt+1:])

	return &TreeNode{nums[maxAt], left, right}
}
