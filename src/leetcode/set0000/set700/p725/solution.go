package p725

import (
	"bytes"
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
	var buf bytes.Buffer
	buf.WriteString("[")

	for tmp := list; tmp != nil; tmp = tmp.Next {
		buf.WriteString(strconv.Itoa(tmp.Val))
		buf.WriteString(",")
	}

	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}

	buf.WriteString("]")
	return buf.String()
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
	for i := 1; i < len(nodes); i++ {
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	sz := size(root)

	a, b := sz/k, sz%k

	res := make([]*ListNode, k)

	head := root
	for i := 0; i < k && head != nil; i++ {
		res[i] = head
		tmp := head
		for j := 0; j < a; j++ {
			tmp = head
			head = head.Next
		}
		if b > 0 {
			tmp = head
			head = head.Next
			b--
		}
		tmp.Next = nil
	}
	return res
}

func size(root *ListNode) int {
	res := 0
	tmp := root
	for tmp != nil {
		res++
		tmp = tmp.Next
	}
	return res
}
