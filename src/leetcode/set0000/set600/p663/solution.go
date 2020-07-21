package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

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

func checkEqualTree(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return false
	}

	var sum func(node *TreeNode) int
	cache := make(map[*TreeNode]int)

	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		tmp := node.Val + sum(node.Left) + sum(node.Right)
		cache[node] = tmp
		return tmp
	}

	total := sum(root)

	if total%2 != 0 {
		return false
	}

	half := total / 2

	for _, tmp := range cache {
		if tmp == half {
			return true
		}
	}

	return false
}

func main() {
	root := ParseAsTree("[-9,-3,2,null,4,4,0,-6,null,-5]")
	fmt.Println(checkEqualTree(root))
}
