package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{2, &ListNode{1, nil}}
	head = partition(head, 2)
	for a := head; a != nil; a = a.Next {
		fmt.Println(a.Val)
	}
}

func partition(head *ListNode, x int) *ListNode {
	a := &list{}
	b := &list{}

	c := head
	for c != nil {
		if c.Val >= x {
			a.add(c)
		} else {
			b.add(c)
		}
		d := c.Next
		c.Next = nil
		c = d
	}

	return concat(b, a)
}

type list struct {
	head *ListNode
	tail *ListNode
}

func (l *list) add(node *ListNode) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	l.tail.Next = node
	l.tail = l.tail.Next
}

func concat(a, b *list) *ListNode {
	if a.head == nil {
		return b.head
	}

	c := b.head

	for c != nil {
		a.tail.Next = c
		a.tail = a.tail.Next
		c = c.Next
	}

	return a.head
}
