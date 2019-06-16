package p979

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

func distributeCoins(root *TreeNode) int {

	var res int

	var loop func(node *TreeNode)

	loop = func(node *TreeNode) {
		if node == nil {
			return
		}

		lc := countNodes(node.Left)
		ls := sumNodes(node.Left)
		rc := countNodes(node.Right)
		rs := sumNodes(node.Right)

		if lc < ls {
			//left has more, need to pull up some values from left children
			res += ls - lc
			node.Val += ls - lc
		} else if lc > ls {
			//need to push down some value from root
			res += lc - ls
			node.Val -= lc - ls
		}

		//same for right
		if rc < rs {
			res += rs - rc
			node.Val += rs - rc
		} else if rc > rs {
			res += rc - rs
			node.Val -= rc - rs
		}

		loop(node.Left)
		loop(node.Right)
	}

	loop(root)

	return res

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func sumNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sumNodes(root.Left) + sumNodes(root.Right)
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
