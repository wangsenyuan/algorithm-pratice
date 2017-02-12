package main

import (
	"math"
	"fmt"
)

func main() {
	root := &TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, nil}, &TreeNode{2, nil, &TreeNode{9, nil, nil}}}
	fmt.Println(findValueMostElement(root))
}

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findValueMostElement(root *TreeNode) []int {

	var height func(root *TreeNode) int

	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(height(root.Left), height(root.Right)) + 1
	}

	h := height(root)
	if h == 0 {
		return nil
	}

	res := make([]int, h)

	for i := 0; i < h; i++ {
		res[i] = math.MinInt32
	}

	var visit func(root *TreeNode, h int)

	visit = func(root *TreeNode, h int) {
		if root == nil {
			return
		}
		if root.Val > res[h] {
			res[h] = root.Val
		}
		visit(root.Left, h+1)
		visit(root.Right, h+1)
	}

	visit(root, 0)

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
