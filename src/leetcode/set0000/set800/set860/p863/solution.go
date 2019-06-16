package p863

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseStrAsTree(str string) *TreeNode {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return nil
	}
	if str[0] == '[' && str[len(str)-1] == ']' {
		str = str[1 : len(str)-1]
	}
	ss := strings.Split(str, ",")
	trees := make([]*TreeNode, len(ss))
	for i := 0; i < len(ss); i++ {
		if ss[i] == "null" {
			continue
		}
		x, _ := strconv.Atoi(ss[i])
		node := new(TreeNode)
		node.Val = x
		if i > 0 {
			parent := trees[(i-1)/2]
			if i%2 == 1 {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}
		trees[i] = node
	}

	return trees[0]
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {

	var dfs func(node *TreeNode) int

	var dfs2 func(node *TreeNode, dist int)

	res := make([]int, 0, 10)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		if node == target || node.Val == target.Val {
			dfs2(node, K)
			return 0
		}
		ld := dfs(node.Left)
		if ld >= 0 {
			ld++
			if ld == K {
				res = append(res, node.Val)
			} else {
				dfs2(node.Right, K-ld-1)
			}

			return ld
		}
		rd := dfs(node.Right)

		if rd >= 0 {
			rd++
			if rd == K {
				res = append(res, node.Val)
			} else {
				dfs2(node.Left, K-rd-1)
			}

			return rd
		}
		return -1
	}

	dfs2 = func(node *TreeNode, dist int) {
		if node == nil || dist < 0 {
			return
		}
		if dist == 0 {
			res = append(res, node.Val)
			return
		}
		dfs2(node.Left, dist-1)
		dfs2(node.Right, dist-1)
	}

	dfs(root)

	return res
}
