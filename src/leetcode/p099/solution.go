package main

import "fmt"

func main() {
	// root := &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}
	root := &TreeNode{0, &TreeNode{1, nil, nil}, nil}
	recoverTree(root)
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

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev *TreeNode
	var first *TreeNode
	var second *TreeNode
	current := root

	for current != nil {
		if current.Left != nil {
			tmp := current.Left
			for tmp.Right != nil && tmp.Right != current {
				tmp = tmp.Right
			}

			if tmp.Right == nil {
				tmp.Right = current
				current = current.Left
			} else {
				first, second = findSwappedNodes(first, second, prev, current)
				prev = current
				tmp.Right = nil
				current = current.Right
			}
		} else {
			first, second = findSwappedNodes(first, second, prev, current)
			prev = current
			current = current.Right
		}
	}

	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func findSwappedNodes(first, second, prev, current *TreeNode) (*TreeNode, *TreeNode) {
	if prev != nil && prev.Val > current.Val {
		if first == nil {
			return prev, current
		}

		return first, current
	}

	return first, second
}
