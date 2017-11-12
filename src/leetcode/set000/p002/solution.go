package main

import "fmt"

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	if node.Next == nil {
		return fmt.Sprintf("%d", node.Val)
	}

	return fmt.Sprintf("%d(%v)", node.Val, node.Next.String())
}

func next(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}

func value(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var add func(a *ListNode, b *ListNode, carray int) *ListNode

	add = func(a *ListNode, b *ListNode, carray int) *ListNode {
		if a == nil && b == nil && carray == 0 {
			return nil
		}

		sum := value(a) + value(b) + carray

		if sum < 10 {
			return &ListNode{sum, add(next(a), next(b), 0)}
		}

		return &ListNode{sum - 10, add(next(a), next(b), 1)}
	}

	return add(l1, l2, 0)
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("%v\n", l3)
}
