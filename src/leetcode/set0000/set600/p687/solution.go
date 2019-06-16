package main

import (
	. "../util"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var visit func(node *TreeNode) (int, int)

	visit = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		a, b := visit(node.Left)
		c, d := visit(node.Right)

		if node.Left != nil && node.Left.Val == node.Val && node.Right != nil && node.Right.Val == node.Val {
			e := b + d + 1
			return max(a, max(c, e)), max(b, d) + 1
		}

		if node.Left != nil && node.Left.Val == node.Val {
			e := b + 1
			return max(a, max(c, e)), e
		}

		if node.Right != nil && node.Right.Val == node.Val {
			e := d + 1
			return max(a, max(c, e)), e
		}

		return max(a, max(c, 1)), 1
	}

	ans, _ := visit(root)
	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tree := ParseAsTree("[5,4,5,1,1,null,5]")
	fmt.Println(longestUnivaluePath(tree))
	tree = ParseAsTree("[1,4,5,4,4,null,5]")
	fmt.Println(longestUnivaluePath(tree))
	tree = ParseAsTree("[1,2,3]")
	fmt.Println(longestUnivaluePath(tree))
}
