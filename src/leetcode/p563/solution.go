package main

import "fmt"

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(findTilt(root))
}

/**
 * Definition for a binary tree node.
  */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var visit func(root *TreeNode) (int64, int64)

	visit = func(root *TreeNode) (int64, int64) {
		if root == nil {
			return 0, 0
		}

		leftSum, leftTilt := visit(root.Left)
		rightSum, rightTilt := visit(root.Right)

		curTilt := abs(leftSum - rightSum)
		return leftSum + rightSum + int64(root.Val), curTilt + leftTilt + rightTilt
	}

	_, tilt := visit(root)
	return int(tilt)
}
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
