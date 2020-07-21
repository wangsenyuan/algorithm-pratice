package main

import (
	"fmt"
	"math"
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
