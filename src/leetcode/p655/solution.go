package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {

	var findWidth func(root *TreeNode) (int, int)

	findWidth = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftWidth, leftHeight := findWidth(root.Left)
		rightWidth, rightHeight := findWidth(root.Right)
		width := max(leftWidth, rightWidth)
		height := max(leftHeight, rightHeight)
		return 2*width + 1, height + 1
	}

	width, height := findWidth(root)

	res := make([][]string, height)
	for i := range res {
		res[i] = make([]string, width)
	}

	var visit func(root *TreeNode, height int, leftWidht int, curWidth int)

	visit = func(root *TreeNode, height int, leftWidth int, curWidth int) {
		if root == nil {
			return
		}
		res[height][leftWidth] = strconv.Itoa(root.Val)
		visit(root.Left, height+1, leftWidth-curWidth/2-1, curWidth/2)
		visit(root.Right, height+1, leftWidth+curWidth/2+1, curWidth/2)
	}

	visit(root, 0, width/2, width/2)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}
	res := printTree(root)
	for _, row := range res {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}

func test2() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, nil}, nil}, &TreeNode{5, nil, nil}}
	res := printTree(root)
	for _, row := range res {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}

func test3() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	res := printTree(root)
	for _, row := range res {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}
