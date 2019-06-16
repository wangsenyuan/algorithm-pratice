package main

import (
	. "../util"
	"fmt"
)

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val >= L && root.Val <= R {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}

	left := trimBST(root.Left, L, R)
	right := trimBST(root.Right, L, R)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Right == nil {
		left.Right = right
		return left
	}

	if right.Left == nil {
		right.Left = left
		return right
	}

	parent, newRoot := left, left.Right

	for parent.Right != nil {
		parent, newRoot = newRoot, parent.Right
	}

	parent.Right = nil
	newRoot.Right = right
	newRoot.Left = left

	return newRoot
}

func main() {
	tree := ParseAsTree("[3,0,4,null,2,null,null,1]")
	tree = trimBST(tree, 1, 3)
	fmt.Printf(SprintTree(tree))
}
