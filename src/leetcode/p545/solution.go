package main

import "fmt"

func main() {
	//root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}}

	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(boundaryOfBinaryTree(root))
}

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var firstLeft, lastLeft, lastLeaf *TreeNode
	var res []int
	var leftBoundary, rightBoundary func(root *TreeNode)

	leftBoundary = func(root *TreeNode) {
		if root == nil {
			return
		}
		if firstLeft == nil {
			firstLeft = root
		}
		lastLeft = root
		res = append(res, root.Val)

		if root.Left != nil {
			leftBoundary(root.Left)
		} else {
			leftBoundary(root.Right)
		}
	}

	rightBoundary = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Right != nil {
			rightBoundary(root.Right)
		} else {
			rightBoundary(root.Left)
		}

		if root != lastLeaf && root != firstLeft {
			res = append(res, root.Val)
		}
	}

	var leafBoundary func(root *TreeNode)
	leafBoundary = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if root != lastLeft {
				res = append(res, root.Val)
			}

			lastLeaf = root
			return
		}
		leafBoundary(root.Left)
		leafBoundary(root.Right)
	}

	if root.Left != nil {
		leftBoundary(root)
	} else {
		res = append(res, root.Val)
		firstLeft = root
		lastLeft = root
	}

	leafBoundary(root)

	if root.Right != nil {
		rightBoundary(root)
	}

	return res
}
