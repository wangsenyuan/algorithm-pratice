package main

import "fmt"

func main() {
	//root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}

	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, &TreeNode{5, &TreeNode{7, nil, nil}, nil}, &TreeNode{6, nil, nil}}}
	fmt.Println(findLeftMostNode(root))
}

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeftMostNode(root *TreeNode) int {

	var find func(root *TreeNode, height int) (int, int)

	find = func(root *TreeNode, height int) (int, int) {
		if root == nil {
			return -1, height - 1
		}

		if root.Left == nil && root.Right == nil {
			return root.Val, height
		}

		lv, lh := find(root.Left, height+1)
		rv, rh := find(root.Right, height+1)

		if lh >= rh {
			return lv, lh
		}
		return rv, rh
	}

	v, _ := find(root, 0)

	return v
}
