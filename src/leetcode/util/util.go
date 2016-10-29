package util

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) String() string {
	if list == nil {
		return "nil"
	}
	return fmt.Sprintf("%d -> %v", list.Val, list.Next)
}

func ParseAsList(str string) *ListNode {
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	nodes := strings.Split(str, ",")

	if len(nodes) == 0 {
		return nil
	}

	head := &ListNode{Val: parseNum(nodes[0])}
	prev := head
	for i := range nodes {
		cur := &ListNode{Val: parseNum(nodes[i])}
		prev.Next = cur
		prev = cur
	}

	return head
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
