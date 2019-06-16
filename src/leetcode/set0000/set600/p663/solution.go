package main

import (
	. "../util"
	"fmt"
)

/**
 * Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func checkEqualTree(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return false
	}

	var sum func(node *TreeNode) int
	cache := make(map[*TreeNode]int)

	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		tmp := node.Val + sum(node.Left) + sum(node.Right)
		cache[node] = tmp
		return tmp
	}

	total := sum(root)

	if total%2 != 0 {
		return false
	}

	half := total / 2

	for _, tmp := range cache {
		if tmp == half {
			return true
		}
	}

	return false
}

func main() {
	root := ParseAsTree("[-9,-3,2,null,4,4,0,-6,null,-5]")
	fmt.Println(checkEqualTree(root))
}
