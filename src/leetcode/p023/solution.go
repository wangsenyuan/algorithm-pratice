package main

import "math"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func value(node *ListNode) int {
	if node != nil {
		return node.Val
	}
	return math.MaxInt32
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}
	return nil
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	return merge(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func merge(a, b *ListNode) *ListNode {
	if a == nil && b == nil {
		return nil
	}

	if value(a) <= value(b) {
		return &ListNode{value(a), merge(next(a), b)}
	}
	return &ListNode{value(b), merge(a, next(b))}
}
