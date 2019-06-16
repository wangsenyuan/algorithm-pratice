package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, nil}
	res := rightSideView(root)
	fmt.Println(res)
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	visit(root, &res, 0, -1)
	return res
}

func visit(node *TreeNode, result *[]int, level int, gate int) int {
	if node == nil {
		if level-1 > gate {
			return level - 1
		}
		return gate
	}

	if level > gate {
		*result = append(*result, node.Val)
	}

	right := visit(node.Right, result, level+1, gate)
	left := visit(node.Left, result, level+1, right)

	if left > right {
		return left
	}
	return right
}
