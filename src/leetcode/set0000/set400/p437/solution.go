package main

import (
	"fmt"

	. "../util"
)

func main() {
	root := ParseAsTree("1,-2,-3,1,3,-2,null,-1")
	fmt.Println(pathSum(root, -1))
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode, cur int)
	cnt := 0
	dfs = func(root *TreeNode, cur int) {
		if root == nil {
			return
		}
		cur += root.Val
		if cur == sum {
			cnt++
		}

		dfs(root.Left, cur)
		dfs(root.Right, cur)
	}

	dfs(root, 0)
	b := pathSum(root.Left, sum)
	c := pathSum(root.Right, sum)

	return cnt + b + c
}
