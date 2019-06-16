package main

import "fmt"

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	res := preorderTraversal(root)
	fmt.Printf("%v\n", res)
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int

	for root != nil {
		if root.Left != nil {
			tmp := root.Left
			for tmp.Right != nil && tmp.Right != root {
				tmp = tmp.Right
			}

			if tmp.Right == nil {
				res = append(res, root.Val)
				tmp.Right = root
				root = root.Left
			} else {
				tmp.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}
