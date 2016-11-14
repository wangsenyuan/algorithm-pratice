package main

import (
	"fmt"

	. "../util"
)

func main() {
	root := ParseAsTree("[3,2,3,#,3,#,1]")
	fmt.Println(rob(root))
}

func rob(root *TreeNode) int {

	var doRob func(root *TreeNode) (int, int)

	doRob = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		robLeft, skipLeft := doRob(root.Left)
		robRight, skipRight := doRob(root.Right)

		robThis := max(skipLeft+skipRight+root.Val, robLeft+robRight)
		return robThis, robLeft + robRight
	}

	return max(doRob(root))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
