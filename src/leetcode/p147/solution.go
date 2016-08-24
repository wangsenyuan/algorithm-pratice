package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{0, &ListNode{2, nil}}}
	head = insertionSortList(head)
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

func insertionSortList(head *ListNode) *ListNode {
	var prev *ListNode
	a := head

	for a != nil {
		b := a.Next

		if prev == nil || prev.Val >= a.Val {
			a.Next = prev
			prev = a
		} else {
			c := prev
			for c.Next != nil && c.Next.Val < a.Val {
				c = c.Next
			}

			d := c.Next
			c.Next = a
			a.Next = d
		}

		a = b
	}

	return prev
}
