package main

import "fmt"

/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	return fmt.Sprintf("%d", ln.Val)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	newHead := &ListNode{-1, head}
	prev := newHead
	curr := head
	var next *ListNode
	for i := 1; i < n; i++ {
		if i < m {
			prev = curr
			curr = curr.Next
		} else {
			next = curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}

	}

	return newHead.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = reverseBetween(head, 2, 4)
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}
