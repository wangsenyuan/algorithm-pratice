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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	levels := []int{0}
	var result [][]int
	lastLevel := 0
	var lastRow []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		level := levels[0]
		levels = levels[1:]
		if level > lastLevel {
			lastLevel = level
			result = append(result, lastRow)
			lastRow = make([]int, 0)
		}
		lastRow = append(lastRow, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
			levels = append(levels, level+1)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			levels = append(levels, level+1)
		}
	}
	result = append(result, lastRow)
	return result
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	result := levelOrder(root)
	for _, row := range result {
		for _, x := range row {
			fmt.Printf("%d ", x)
		}
		fmt.Println()
	}
}
