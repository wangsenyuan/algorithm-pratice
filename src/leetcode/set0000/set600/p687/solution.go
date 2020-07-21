package main

import (
	"fmt"
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var visit func(node *TreeNode) (int, int)

	visit = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		a, b := visit(node.Left)
		c, d := visit(node.Right)

		if node.Left != nil && node.Left.Val == node.Val && node.Right != nil && node.Right.Val == node.Val {
			e := b + d + 1
			return max(a, max(c, e)), max(b, d) + 1
		}

		if node.Left != nil && node.Left.Val == node.Val {
			e := b + 1
			return max(a, max(c, e)), e
		}

		if node.Right != nil && node.Right.Val == node.Val {
			e := d + 1
			return max(a, max(c, e)), e
		}

		return max(a, max(c, 1)), 1
	}

	ans, _ := visit(root)
	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tree := ParseAsTree("[5,4,5,1,1,null,5]")
	fmt.Println(longestUnivaluePath(tree))
	tree = ParseAsTree("[1,4,5,4,4,null,5]")
	fmt.Println(longestUnivaluePath(tree))
	tree = ParseAsTree("[1,2,3]")
	fmt.Println(longestUnivaluePath(tree))
}
