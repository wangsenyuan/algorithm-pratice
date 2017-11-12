package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{2, nil, nil}, nil}}
	fmt.Println(findMode(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {

	var find func(root *TreeNode, cnts map[int]int)
	find = func(root *TreeNode, cnts map[int]int) {
		if root == nil {
			return
		}
		cnts[root.Val]++
		find(root.Left, cnts)
		find(root.Right, cnts)

	}

	cnts := make(map[int]int)
	find(root, cnts)

	mx := 0
	for _, cnt := range cnts {
		if cnt > mx {
			mx = cnt
		}
	}

	var ans []int

	for val, cnt := range cnts {
		if cnt == mx {
			ans = append(ans, val)
		}
	}

	return ans
}
