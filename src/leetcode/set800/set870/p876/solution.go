package p876

import (
	"fmt"
	"os"
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

func middleNode(head *ListNode) *ListNode {
	fmt.Fprintf(os.Stderr, "%s\n", head.String())
	if head == nil {
		return nil
	}
	cnt := size(head)
	mid := cnt / 2

	node := head

	for i := 0; i < mid; i++ {
		node = node.Next
	}

	return node
}

func size(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + size(head.Next)
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
	for i := 1; i < len(nodes); i++ {
		prev.Next = &ListNode{Val: parseNum(nodes[i])}
		prev = prev.Next
	}

	return head
}

func parseNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
