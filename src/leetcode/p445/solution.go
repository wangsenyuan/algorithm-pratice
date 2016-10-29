package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l1 := ParseAsList("[7,2,4,3]")
	l2 := ParseAsList("[5,6,4]")
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3)
}

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
		prev.Next = &ListNode{Val: parseNum(nodes[i])}
		prev = prev.Next
	}

	return head
}

func parseNum(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var add0 func(a, b *ListNode) (*ListNode, int)
	var add1 func(a, b, l3 *ListNode) (*ListNode, int)
	var add2 func(a, b *ListNode) (*ListNode, int)

	add0 = func(a, b *ListNode) (*ListNode, int) {
		if a.Next == nil && b.Next == nil {
			return add2(l1, l2)
		}
		if a.Next == nil {
			return add1(l2, b.Next, l1)
		}

		if b.Next == nil {
			return add1(l1, a.Next, l2)
		}

		return add0(a.Next, b.Next)
	}

	add1 = func(a, b, l3 *ListNode) (*ListNode, int) {
		var next *ListNode
		carry := 0
		if b.Next != nil {
			next, carry = add1(a.Next, b.Next, l3)
		} else {
			next, carry = add2(a.Next, l3)
		}
		sum := a.Val + carry
		if sum >= 10 {
			return &ListNode{sum - 10, next}, 1
		}
		return &ListNode{sum, next}, 0
	}

	add2 = func(a, b *ListNode) (*ListNode, int) {
		if a == nil {
			return nil, 0
		}
		next, carry := add2(a.Next, b.Next)
		sum := a.Val + b.Val + carry
		if sum >= 10 {
			return &ListNode{sum - 10, next}, 1
		}
		return &ListNode{sum, next}, 0
	}

	l3, carry := add0(l1, l2)
	if carry == 0 {
		return l3
	}

	return &ListNode{1, l3}
}
