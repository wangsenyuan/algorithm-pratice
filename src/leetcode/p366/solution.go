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

func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0, 10)

	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		d := max(left, right) + 1
		if len(result) == d {
			result = append(result, make([]int, 0, 10))
		}
		result[d] = append(result[d], node.Val)
		return d
	}

	dfs(root)

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	leaves := findLeaves(root)
	for _, row := range leaves {
		fmt.Printf("%v\n", row)
	}
	root = &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}
	leaves = findLeaves(root)
	for _, row := range leaves {
		fmt.Printf("%v\n", row)
	}
}
