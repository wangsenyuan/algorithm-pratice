package p867

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

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	var depth func(node *TreeNode, d int)

	var maxCnt int
	var maxDepth int
	depth = func(node *TreeNode, d int) {
		if node == nil {
			return
		}

		if d > maxDepth {
			maxDepth = d
			maxCnt = 1
		} else if d == maxDepth {
			maxCnt++
		}

		depth(node.Left, d+1)
		depth(node.Right, d+1)
	}

	depth(root, 0)

	var dfs func(node *TreeNode, d int) int
	var ans *TreeNode
	dfs = func(node *TreeNode, d int) int {
		if node == nil {
			return 0
		}
		if d == maxDepth {
			if maxCnt == 1 {
				ans = node
			}
			return 1
		}
		cnt := dfs(node.Left, d+1) + dfs(node.Right, d+1)
		if cnt == maxCnt {
			if ans == nil {
				ans = node
			}
		}
		return cnt
	}
	dfs(root, 0)
	return ans
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
