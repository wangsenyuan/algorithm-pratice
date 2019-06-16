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

func tree2str(t *TreeNode) string {

	var visit func(t *TreeNode) string

	visit = func(t *TreeNode) string {
		if t == nil {
			return ""
		}

		if t.Left == nil && t.Right == nil {
			return fmt.Sprintf("%d", t.Val)
		}

		res := fmt.Sprintf("%d", t.Val)
		res += "(" + visit(t.Left) + ")"

		if t.Right != nil {
			res += "(" + visit(t.Right) + ")"
		}

		return res
	}

	return visit(t)

}

func main() {
	t := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(tree2str(t))
}
