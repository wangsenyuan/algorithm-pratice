package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"fmt"

	. "../util"
)

func main() {
	root := ParseAsTree("[2,1,3,null,null,null,4]")
	deleteNode(root, 3)
	fmt.Println(root)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Right != nil {
			tmp := root.Right
			for tmp.Left != nil {
				tmp = tmp.Left
			}
			root.Val = tmp.Val
			root.Right = deleteNode(root.Right, tmp.Val)
			return root
		}
		return root.Left
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
