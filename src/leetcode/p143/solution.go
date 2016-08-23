package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head := &ListNode{1, nil}
	reorderList(head)

	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	a := head
	n := 0
	var prev *ListNode

	for a != nil {
		n++
		b := a.Next
		a.Next = prev
		prev = a
		a = b
	}

	mid := n / 2

	a = prev

	x := a
	for i := 0; i < mid - 1; i++ {
		x = x.Next
	}

	prev = x.Next
	x.Next = nil
	x = prev
	prev = nil

	for x != nil {
		b := x.Next
		x.Next = prev
		prev = x
		x = b
	}

	head = prev
	x = prev
	for i := 0; i < mid; i++ {
		d := x.Next
		c := a.Next
		x.Next = a
		x = d
		a.Next = x
		a = c
	}
}
