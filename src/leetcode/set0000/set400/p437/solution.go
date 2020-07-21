package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	root := ParseAsTree("1,-2,-3,1,3,-2,null,-1")
	fmt.Println(pathSum(root, -1))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	if tn == nil {
		return "null"
	}
	return fmt.Sprintf("%d,%v,%v", tn.Val, tn.Left, tn.Right)
}

func ParseAsTree(str string) *TreeNode {
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)

	var currLevel []*TreeNode
	var nextLevel []*TreeNode
	nodes := strings.Split(str, ",")

	root := parseNode(nodes[0])
	currLevel = append(currLevel, root)

	pt := 0
	ct := 0
	for i := 1; i < len(nodes); i++ {
		parent := currLevel[pt]
		node := parseNode(nodes[i])
		if node != nil {
			nextLevel = append(nextLevel, node)
		}
		if ct == 0 {
			parent.Left = node
			ct++
		} else {
			parent.Right = node
			pt++
			ct = 0
		}
		if pt == len(currLevel) {
			currLevel = nextLevel
			nextLevel = make([]*TreeNode, 0)
			pt = 0
		}
	}
	return root
}

func parseNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	return &TreeNode{parseNum(str), nil, nil}
}

func parseNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
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
