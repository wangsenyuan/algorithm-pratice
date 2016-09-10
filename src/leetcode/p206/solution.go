package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	head = reverseList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	a := head
	b := a.Next
	a.Next = nil

	for b != nil {
		c := b.Next
		b.Next = a
		a, b = b, c
	}

	return a
}
