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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	count := make(map[string]int)

	var visit func(root *TreeNode) string

	visit = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		str := fmt.Sprintf("%d(%s)(%s)", root.Val, visit(root.Left), visit(root.Right))

		if cnt, found := count[str]; found {
			if cnt == 1 {
				res = append(res, root)
			}
			count[str] = cnt + 1
		} else {
			count[str] = 1
		}
		return str
	}

	visit(root)

	return res
}

func main() {
	root := &TreeNode{2, &TreeNode{1, &TreeNode{2, nil, nil}, nil}, &TreeNode{3, &TreeNode{1, &TreeNode{2, nil, nil}, nil}, nil}}
	res := findDuplicateSubtrees(root)
	for _, node := range res {
		fmt.Println(node.Val)
	}
}
