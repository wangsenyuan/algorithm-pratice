package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{0, &ListNode{2, &ListNode{1, nil}}}}
	head = sortList(head)
	for head != nil {
		fmt.Printf("%d->", head.Val)
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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil

	return merge(sortList(head), sortList(mid))
}

func merge(a, b *ListNode) *ListNode {
	head := &ListNode{0, nil}
	prev := head
	for a != nil || b != nil {
		if b == nil || (a != nil && a.Val <= b.Val) {
			prev.Next = a
			prev = a
			a = a.Next
		} else {
			prev.Next = b
			prev = b
			b = b.Next
		}
	}

	return head.Next
}
