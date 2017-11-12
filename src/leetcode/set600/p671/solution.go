package main

import (
	. "../util"
	"math"
	"fmt"
)

func findSecondMinimumValue(root *TreeNode) int {
	minmum := root.Val

	var visit func(node *TreeNode) int

	visit = func(node *TreeNode) int {
		if node == nil {
			return math.MaxInt32
		}

		if node.Val > minmum {
			return node.Val
		}
		left := visit(node.Left)
		right := visit(node.Right)
		if left < right {
			return left
		}
		return right
	}

	ans := visit(root)

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	tree := ParseAsTree("[2,2,2]")
	fmt.Println(findSecondMinimumValue(tree))
	tree = ParseAsTree("[2,2,5,null,null,5,7]")
	fmt.Println(findSecondMinimumValue(tree))
}
